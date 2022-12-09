package main

import (
	"github.com/zhangel/logger"
	"tip/tools/import_db/core"

	//"tip/tools/import_db/core"
	"github.com/go-ini/ini"
)

/*
	var src string
	var dest string
	rootCmd:=&cobra.Command{
		Use: "import_db",
		Short: "short",
		Long: "long explain",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	sysinfoCmd:=&cobra.Command{
		Use: "sysinfo",
		Short:"short",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(src)
			fmt.Println(dest)
		},
	}
	rootCmd.AddCommand(sysinfoCmd)
	sysinfoCmd.Flags().StringVarP(&dest,"dest","d","","dest path")
	sysinfoCmd.Flags().StringVarP(&src,"src","s","","src path")
	rootCmd.Execute()
*/

func String(key string) string {
	configPath:="./config.ini"
	cfg,err:=ini.Load(configPath)
	if err != nil {
		logger.Fatal("load config file fail,error=%+v",err)
	}
	s:=cfg.Section("")
	return s.Key(key).String()
}

func main() {
	core.NewCore().Run()
}
