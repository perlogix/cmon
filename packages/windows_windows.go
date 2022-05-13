//go:build windows

package packages

import (
	"runtime"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/osconfig/packages"
	"github.com/perlogix/cmon/data"
	"golang.org/x/sys/windows/registry"
)

func parseDate(dateString string) time.Time {
	if len(dateString) != 8 {
		return time.Time{}
	}

	year, err := strconv.ParseInt(dateString[0:4], 10, 32)
	if err != nil {
		return time.Time{}
	}
	month, err := strconv.ParseInt(dateString[4:6], 10, 32)
	if err != nil {
		return time.Time{}
	}
	day, err := strconv.ParseInt(dateString[6:8], 10, 32)
	if err != nil {
		return time.Time{}
	}

	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Now().Location())
}

func getWindowsApplication(k *registry.Key) *packages.WindowsApplication {
	displayName, _, errName := k.GetStringValue("DisplayName")
	_, _, errUninstall := k.GetStringValue("UninstallString")

	if errName == nil && errUninstall == nil {
		displayVersion, _, _ := k.GetStringValue("DisplayVersion")
		publisher, _, _ := k.GetStringValue("Publisher")
		installDate, _, _ := k.GetStringValue("InstallDate")
		helpLink, _, _ := k.GetStringValue("HelpLink")
		return &packages.WindowsApplication{
			DisplayName:    displayName,
			DisplayVersion: displayVersion,
			Publisher:      publisher,
			InstallDate:    parseDate(installDate),
			HelpLink:       helpLink,
		}
	}
	return nil
}

func GetWindowsApplications() ([]*packages.WindowsApplication, error) {
	directories := []string{
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
		`SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
	}
	var allApps []*packages.WindowsApplication

	for _, dir := range directories {
		apps, err := getWindowsApplications(dir)
		if err != nil {
			continue
		}
		allApps = append(allApps, apps...)
	}
	return allApps, nil
}

func getWindowsApplications(directory string) ([]*packages.WindowsApplication, error) {
	dirKey, err := registry.OpenKey(registry.LOCAL_MACHINE, directory, registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		return nil, err
	}
	defer dirKey.Close()

	var result []*packages.WindowsApplication
	subkeys, err := dirKey.ReadSubKeyNames(0)
	if err != nil {
		return nil, err
	}
	for _, subkey := range subkeys {
		k, err := registry.OpenKey(dirKey, subkey, registry.QUERY_VALUE)
		if err != nil {
			continue
		}
		app := getWindowsApplication(&k)
		if app != nil {
			result = append(result, app)
		}
		k.Close()
	}
	return result, nil
}

func WindowsPackages(d *data.DiscoverJSON) {

	packages := []data.WindowsPackages{}

	if runtime.GOOS == "windows" {

		pkgs, err := GetWindowsApplications()
		if err != nil {
			d.WindowsPackages = packages
			return
		}

		wp := data.WindowsPackages{}

		for _, e := range pkgs {
			wp.DisplayName = e.DisplayName
			wp.DisplayVersion = e.DisplayVersion
			wp.InstallDate = e.InstallDate
			wp.Publisher = e.Publisher
			packages = append(packages, wp)
		}
	}

	d.WindowsPackages = packages
}
