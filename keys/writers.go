package keys

import (
	"io"
	//"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/eris-ltd/eris-cli/config"
	"github.com/eris-ltd/eris-cli/data"
	"github.com/eris-ltd/eris-cli/definitions"
	srv "github.com/eris-ltd/eris-cli/services"

	"github.com/eris-ltd/common/go/common"
)

func GenerateKey(do *definitions.Do) error {
	do.Name = "keys"

	if err := srv.EnsureRunning(do); err != nil {
		return err
	}

	buf, err := srv.ExecHandler(do.Name, []string{"eris-keys", "gen", "--no-pass"})
	if err != nil {
		return err
	}

	io.Copy(config.GlobalConfig.Writer, buf)

	return nil
}

func GetPubKey(do *definitions.Do) error {
	do.Name = "keys"
	if err := srv.EnsureRunning(do); err != nil {
		return err
	}

	buf, err := srv.ExecHandler(do.Name, []string{"eris-keys", "pub", "--addr", do.Address})
	if err != nil {
		return err
	}

	io.Copy(config.GlobalConfig.Writer, buf)

	return nil
}

func ExportKey(do *definitions.Do) error {
	do.Name = "keys"
	if err := srv.EnsureRunning(do); err != nil {
		return err
	}

	do.Destination = common.KeysPath // not KeysDataPath because we are exporting the directory
	if do.All && do.Address == "" {
		doLs := definitions.NowDo()
		doLs.Container = true
		doLs.Host = false
		doLs.Quiet = true
		if err := ListKeys(doLs); err != nil {
			return err
		}
		keyArray := strings.Split(do.Result, ",")

		for _, addr := range keyArray {
			do.Source = path.Join(common.ErisContainerRoot, "keys", "data", addr)
			if err := data.ExportData(do); err != nil {
				return err
			}
		}
	} else {
		do.Source = path.Join(common.ErisContainerRoot, "keys", "data", do.Address)
		if err := data.ExportData(do); err != nil {
			return err
		}
	}
	return nil
}

func ImportKey(do *definitions.Do) error {
	do.Name = "keys"
	if err := srv.EnsureRunning(do); err != nil {
		return err
	}

	do.Destination = path.Join(common.KeysContainerPath, do.Address)
	if do.All && do.Address == "" {
		doLs := definitions.NowDo()
		doLs.Container = false
		doLs.Host = true
		doLs.Quiet = true
		if err := ListKeys(doLs); err != nil {
			return err
		}
		keyArray := strings.Split(do.Result, ",")

		for _, addr := range keyArray {
			do.Source = filepath.Join(common.KeysPath, "data", addr)
			do.Destination = path.Join(common.ErisContainerRoot, "keys", "data", addr)
			if err := data.ImportData(do); err != nil {
				return err
			}
		}
	} else {
		do.Source = filepath.Join(common.KeysDataPath, do.Address, do.Address)
		if err := data.ImportData(do); err != nil {
			return err
		}
	}

	return nil
}

func ConvertKey(do *definitions.Do) error {
	do.Name = "keys"
	if err := srv.EnsureRunning(do); err != nil {
		return err
	}

	buf, err := srv.ExecHandler(do.Name, []string{"mintkey", "mint", do.Address})
	if err != nil {
		return err
	}

	io.Copy(config.GlobalConfig.Writer, buf)

	return nil
}
