#!/bin/sh

# Main variables
START=$(date +%s)
LOG="build.log"
PROJECT=$(basename "$(pwd)")
VERSION=$(git log -n 1 --pretty=format:"%H" 2>/dev/null | head -c 10)

if [ "$VERSION" = "" ]; then
    # Version will return YYYYmmdd, e.g. 20210209
    VERSION=$(date '+%Y%m%d')
fi

IMAGE="$PROJECT:$VERSION"

# HTTP proxy if applicable
http_proxy="$http_proxy"
https_proxy="$https_proxy"
HTTP_PROXY="$HTTP_PROXY"
HTTPS_PROXY="$HTTPS_PROXY"
no_proxy="$no_proxy"
NO_PROXY="$NO_PROXY"

# Slack if applicable
SLACK_TOKEN="$SLACK_TOKEN"
SLACK_CHANNELS="$SLACK_CHANNELS"
BUILD_STATUS="PASS"

# Check root permissions
root() {
    if [ "$(id -u)" != 0 ]; then
        echo "Need root permissions" && echo
        exit 1
    fi
}

# Calculate finish time
finish() {
    FINISH=$(date +%s)
    COMPLETE=$((FINISH - START))
    FINISH="took $((COMPLETE / 60)) minute(s) to finish"
}

# Result builder
result() {
    ERRORS=$(grep 'ERRORS FOUND in' "$LOG" | awk '{ print $7 }' | tr '\n' ' ')
    VULNS=$(grep 'Total:' "$LOG")

    if [ "$(grep -i "non-zero code\|error response from daemon" "$LOG" 2>/dev/null)" != "" ]; then
        BUILD_STATUS="FAIL"
    fi

    if [ "$VULNS" != "" ]; then
        RESULT="Vulnerabilities $VULNS"
    else
        VULNS=""
    fi

    finish

    RESULT="$BUILD_STATUS - $PROJECT - $VERSION - $FINISH
    $ERRORS
    $VULNS
    "
}

# Make self-signed certificate
mkcert() {
    openssl req -x509 -newkey \
        rsa:4096 -nodes -keyout localhost.key \
        -out localhost.pem -days 365 -sha256 -subj '/CN=localhost'
}

# Build commands
build() {
    docker build \
        --build-arg http_proxy="$http_proxy" \
        --build-arg https_proxy="$https_proxy" \
        --build-arg HTTP_PROXY="$HTTP_PROXY" \
        --build-arg HTTPS_PROXY="$HTTPS_PROXY" \
        --build-arg no_proxy="$no_proxy" \
        --build-arg NO_PROXY="$NO_PROXY" \
        --network=host -t "$IMAGE" . 2>&1 | tee "$LOG"

    if [ "$(command -v docker-slim)" != "" ]; then
        docker-slim build --http-probe-off \
            --continue-after 5 --tag "$IMAGE" \
            --target "$IMAGE" 2>&1 | tee -a "$LOG"

        docker rm -f "$(docker ps | grep dockerslim | awk '{ print $1 }')"
        rm -f slim.report.json
    fi

    # Remove ASCII colors from log
    sed -ri "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" "$LOG"
}

# Run commands
run() {
    if [ "$(command -v docker-init)" != "" ]; then
        docker rm -f "$PROJECT" 2>/dev/null
        docker run \
            -d --net=host \
            --init --restart=always \
            --name="$PROJECT" "$IMAGE"
    else
        docker rm -f "$PROJECT" 2>/dev/null
        docker run \
            -d --net=host \
            --restart=always \
            --name="$PROJECT" "$IMAGE"
    fi
}

# Clean commands
clean() {
    docker system prune -af
}

# Upload build log file to Slack
slack() {
    if [ "$SLACK_TOKEN" != "" ] && [ -f "$LOG" ]; then
        result

        if [ "$BUILD_STATUS" = "PASS" ]; then
            EMOJI=":white_check_mark:"
        else
            EMOJI=":warning"
        fi

        curl -s \
            -F channels="$SLACK_CHANNELS" \
            -F initial_comment="$EMOJI $RESULT" \
            -F file=@"$LOG" \
            -F filename="$LOG" \
            -H "Authorization: Bearer $SLACK_TOKEN" \
            'https://slack.com/api/files.upload' | grep 'ok":false'
    fi
}

# Check if argument is passed to script
if [ "$1" = "build" ]; then
    root
    build
    exit
fi

if [ "$1" = "run" ]; then
    root
    run
    exit
fi

if [ "$1" = "clean" ]; then
    root
    clean
    exit
fi

if [ "$1" = "mkcert" ]; then
    mkcert
    exit
fi

# Set custom sequence below. Example:
root
build
slack
run
clean

result
cat <<EOF

    $RESULT

EOF
