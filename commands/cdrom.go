package commands

import (
	"context"
	"io"
	"os"

	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	"github.com/ionos-cloud/ionosctl/pkg/resources"
	"github.com/ionos-cloud/ionosctl/pkg/utils"
	"github.com/ionos-cloud/ionosctl/pkg/utils/clierror"
	"github.com/ionos-cloud/ionosctl/pkg/utils/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Server Cdrom Commands

func serverCdrom() *core.Command {
	ctx := context.TODO()
	serverCdromCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "cdrom",
			Short:            "Server CD-ROM Operations",
			Long:             `The sub-commands of ` + "`" + `ionosctl server cdrom` + "`" + ` allow you to attach, get, list, detach CD-ROMs from Servers.`,
			TraverseChildren: true,
		},
	}
	globalFlags := serverCdromCmd.GlobalFlags()
	globalFlags.StringSlice(config.ArgCols, defaultImageCols, "Columns to be printed in the standard output")
	_ = viper.BindPFlag(core.GetGlobalFlagName(serverCdromCmd.Name(), config.ArgCols), globalFlags.Lookup(config.ArgCols))
	_ = serverCdromCmd.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allImageCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Attach Cdrom Command
	*/
	attachCdrom := core.NewCommand(ctx, serverCdromCmd, core.CommandBuilder{
		Namespace: "server",
		Resource:  "cdrom",
		Verb:      "attach",
		ShortDesc: "Attach a CD-ROM to a Server",
		LongDesc: `Use this command to attach a CD-ROM to an existing Server.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option.

Required values to run command:

* Data Center Id
* Server Id
* Cdrom Id`,
		Example:    attachCdromServerExample,
		PreCmdRun:  PreRunDcServerCdromIds,
		CmdRun:     RunServerCdromAttach,
		InitClient: true,
	})
	attachCdrom.AddStringFlag(config.ArgDataCenterId, "", "", config.RequiredFlagDatacenterId)
	_ = attachCdrom.Command.RegisterFlagCompletionFunc(config.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getDataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	attachCdrom.AddStringFlag(config.ArgCdromId, "", "", config.RequiredFlagCdromId)
	_ = attachCdrom.Command.RegisterFlagCompletionFunc(config.ArgCdromId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getImagesCdromIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	attachCdrom.AddStringFlag(config.ArgServerId, "", "", config.RequiredFlagServerId)
	_ = attachCdrom.Command.RegisterFlagCompletionFunc(config.ArgServerId, func(cmd *cobra.Command, ags []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getServersIds(os.Stderr, viper.GetString(core.GetFlagName(attachCdrom.NS, config.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	attachCdrom.AddBoolFlag(config.ArgWaitForRequest, "", config.DefaultWait, "Wait for the Request for CD-ROM attachment to be executed")
	attachCdrom.AddIntFlag(config.ArgTimeout, "", config.DefaultTimeoutSeconds, "Timeout option for Request for Cdrom attachment [seconds]")

	/*
		List Cdroms Command
	*/
	listCdroms := core.NewCommand(ctx, serverCdromCmd, core.CommandBuilder{
		Namespace:  "server",
		Resource:   "cdrom",
		Verb:       "list",
		ShortDesc:  "List attached CD-ROMs from a Server",
		LongDesc:   "Use this command to retrieve a list of CD-ROMs attached to the Server.\n\nRequired values to run command:\n\n* Data Center Id\n* Server Id",
		Example:    listCdromServerExample,
		PreCmdRun:  PreRunDcServerIds,
		CmdRun:     RunServerCdromsList,
		InitClient: true,
	})
	listCdroms.AddStringFlag(config.ArgDataCenterId, "", "", config.RequiredFlagDatacenterId)
	_ = listCdroms.Command.RegisterFlagCompletionFunc(config.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getDataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	listCdroms.AddStringFlag(config.ArgServerId, "", "", config.RequiredFlagServerId)
	_ = listCdroms.Command.RegisterFlagCompletionFunc(config.ArgServerId, func(cmd *cobra.Command, ags []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getServersIds(os.Stderr, viper.GetString(core.GetFlagName(listCdroms.NS, config.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Get Cdrom Command
	*/
	getCdromCmd := core.NewCommand(ctx, serverCdromCmd, core.CommandBuilder{
		Namespace:  "server",
		Resource:   "cdrom",
		Verb:       "get",
		ShortDesc:  "Get a specific attached CD-ROM from a Server",
		LongDesc:   "Use this command to retrieve information about an attached CD-ROM on Server.\n\nRequired values to run command:\n\n* Data Center Id\n* Server Id\n* Cdrom Id",
		Example:    getCdromServerExample,
		InitClient: true,
		PreCmdRun:  PreRunDcServerCdromIds,
		CmdRun:     RunServerCdromGet,
	})
	getCdromCmd.AddStringFlag(config.ArgDataCenterId, "", "", config.RequiredFlagDatacenterId)
	_ = getCdromCmd.Command.RegisterFlagCompletionFunc(config.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getDataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	getCdromCmd.AddStringFlag(config.ArgServerId, "", "", config.RequiredFlagServerId)
	_ = getCdromCmd.Command.RegisterFlagCompletionFunc(config.ArgServerId, func(cmd *cobra.Command, ags []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getServersIds(os.Stderr, viper.GetString(core.GetFlagName(getCdromCmd.NS, config.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	getCdromCmd.AddStringFlag(config.ArgCdromId, "", "", config.RequiredFlagCdromId)
	_ = getCdromCmd.Command.RegisterFlagCompletionFunc(config.ArgCdromId, func(cmd *cobra.Command, ags []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAttachedCdromsIds(os.Stderr, viper.GetString(core.GetFlagName(getCdromCmd.NS, config.ArgDataCenterId)), viper.GetString(core.GetFlagName(getCdromCmd.NS, config.ArgServerId))), cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Detach Cdrom Command
	*/
	detachCdrom := core.NewCommand(ctx, serverCdromCmd, core.CommandBuilder{
		Namespace: "server",
		Resource:  "cdrom",
		Verb:      "detach",
		ShortDesc: "Detach a CD-ROM from a Server",
		LongDesc: `This will detach the CD-ROM from the Server.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option. You can force the command to execute without user input using ` + "`" + `--force` + "`" + ` option.

Required values to run command:

* Data Center Id
* Server Id
* Cdrom Id`,
		Example:    detachCdromServerExample,
		PreCmdRun:  PreRunDcServerCdromIds,
		CmdRun:     RunServerCdromDetach,
		InitClient: true,
	})
	detachCdrom.AddStringFlag(config.ArgDataCenterId, "", "", config.RequiredFlagDatacenterId)
	_ = detachCdrom.Command.RegisterFlagCompletionFunc(config.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getDataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	detachCdrom.AddStringFlag(config.ArgCdromId, "", "", config.RequiredFlagCdromId)
	_ = detachCdrom.Command.RegisterFlagCompletionFunc(config.ArgCdromId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAttachedCdromsIds(os.Stderr, viper.GetString(core.GetFlagName(detachCdrom.NS, config.ArgDataCenterId)),
			viper.GetString(core.GetFlagName(detachCdrom.NS, config.ArgServerId))), cobra.ShellCompDirectiveNoFileComp
	})
	detachCdrom.AddStringFlag(config.ArgServerId, "", "", config.RequiredFlagServerId)
	_ = detachCdrom.Command.RegisterFlagCompletionFunc(config.ArgServerId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getServersIds(os.Stderr, viper.GetString(core.GetFlagName(detachCdrom.NS, config.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	detachCdrom.AddBoolFlag(config.ArgWaitForRequest, "", config.DefaultWait, "Wait for the Request for CD-ROM detachment to be executed")
	detachCdrom.AddIntFlag(config.ArgTimeout, "", config.DefaultTimeoutSeconds, "Timeout option for Request for CD-ROM detachment [seconds]")

	return serverCdromCmd
}

func PreRunDcServerCdromIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.NS, config.ArgDataCenterId, config.ArgServerId, config.ArgCdromId)
}

func RunServerCdromAttach(c *core.CommandConfig) error {
	attachedCdrom, resp, err := c.Servers().AttachCdrom(
		viper.GetString(core.GetFlagName(c.NS, config.ArgDataCenterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgServerId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgCdromId)),
	)
	if err != nil {
		return err
	}

	if err = utils.WaitForRequest(c, printer.GetRequestPath(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getImagePrint(resp, c, getImage(attachedCdrom)))
}

func RunServerCdromsList(c *core.CommandConfig) error {
	attachedCdroms, _, err := c.Servers().ListCdroms(
		viper.GetString(core.GetFlagName(c.NS, config.ArgDataCenterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgServerId)),
	)
	if err != nil {
		return err
	}
	return c.Printer.Print(getImagePrint(nil, c, getCdroms(attachedCdroms)))
}

func RunServerCdromGet(c *core.CommandConfig) error {
	attachedCdrom, _, err := c.Servers().GetCdrom(
		viper.GetString(core.GetFlagName(c.NS, config.ArgDataCenterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgServerId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgCdromId)),
	)
	if err != nil {
		return err
	}
	return c.Printer.Print(getImagePrint(nil, c, getImage(attachedCdrom)))
}

func RunServerCdromDetach(c *core.CommandConfig) error {
	if err := utils.AskForConfirm(c.Stdin, c.Printer, "detach cdrom from server"); err != nil {
		return err
	}
	resp, err := c.Servers().DetachCdrom(
		viper.GetString(core.GetFlagName(c.NS, config.ArgDataCenterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgServerId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgCdromId)),
	)

	if err != nil {
		return err
	}

	if err = utils.WaitForRequest(c, printer.GetRequestPath(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getImagePrint(resp, c, nil))
}

// Output Printing

func getCdroms(cdroms resources.Cdroms) []resources.Image {
	imgs := make([]resources.Image, 0)
	if items, ok := cdroms.GetItemsOk(); ok && items != nil {
		for _, d := range *items {
			imgs = append(imgs, resources.Image{Image: d})
		}
	}
	return imgs
}

func getAttachedCdromsIds(outErr io.Writer, datacenterId, serverId string) []string {
	err := config.Load()
	clierror.CheckError(err, outErr)
	clientSvc, err := resources.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		viper.GetString(config.ArgServerUrl),
	)
	clierror.CheckError(err, outErr)
	serverSvc := resources.NewServerService(clientSvc.Get(), context.TODO())
	cdroms, _, err := serverSvc.ListCdroms(datacenterId, serverId)
	clierror.CheckError(err, outErr)
	attachedCdromsIds := make([]string, 0)
	if items, ok := cdroms.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			if itemId, ok := item.GetIdOk(); ok && itemId != nil {
				attachedCdromsIds = append(attachedCdromsIds, *itemId)
			}
		}
	} else {
		return nil
	}
	return attachedCdromsIds
}

func getImagesCdromIds(outErr io.Writer) []string {
	err := config.Load()
	clierror.CheckError(err, outErr)
	clientSvc, err := resources.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		viper.GetString(config.ArgServerUrl),
	)
	clierror.CheckError(err, outErr)
	imageSvc := resources.NewImageService(clientSvc.Get(), context.TODO())
	images, _, err := imageSvc.List()
	clierror.CheckError(err, outErr)
	imgsIds := make([]string, 0)
	images = sortImagesByType(images, "CDROM")
	if items, ok := images.Images.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			if itemId, ok := item.GetIdOk(); ok && itemId != nil {
				imgsIds = append(imgsIds, *itemId)
			}
		}
	} else {
		return nil
	}
	return imgsIds
}