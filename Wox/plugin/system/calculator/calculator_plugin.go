package calculator

import (
	"context"
	"strings"
	"wox/plugin"
	"wox/share"
	"wox/util"
	"wox/util/clipboard"
)

var calculatorIcon = plugin.NewWoxImageSvg(`<svg t="1697204721503" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5110" width="200" height="200"><path d="M853.333333 341.333333H170.666667v512c0 46.933333 38.4 85.333333 85.333333 85.333334h512c46.933333 0 85.333333-38.4 85.333333-85.333334V341.333333z" fill="#616161" p-id="5111"></path><path d="M768 85.333333H256C209.066667 85.333333 170.666667 123.733333 170.666667 170.666667v192h682.666666V170.666667c0-46.933333-38.4-85.333333-85.333333-85.333334z" fill="#424242" p-id="5112"></path><path d="M768 298.666667H256c-12.8 0-21.333333-8.533333-21.333333-21.333334V170.666667c0-12.8 8.533333-21.333333 21.333333-21.333334h512c12.8 0 21.333333 8.533333 21.333333 21.333334v106.666666c0 12.8-8.533333 21.333333-21.333333 21.333334z" fill="#9CCC65" p-id="5113"></path><path d="M704 213.333333h42.666667v42.666667h-42.666667zM618.666667 213.333333h42.666666v42.666667h-42.666666z" fill="#33691E" p-id="5114"></path><path d="M768 490.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334z" fill="#FF5252" p-id="5115"></path><path d="M320 490.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM469.333333 490.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333334 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333334 21.333334zM618.666667 490.666667h-64c-12.8 0-21.333333-8.533333-21.333334-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333334-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM320 618.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM469.333333 618.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333334 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333334 21.333334zM618.666667 618.666667h-64c-12.8 0-21.333333-8.533333-21.333334-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333334-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM320 746.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM469.333333 746.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333334 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333334 21.333334zM618.666667 746.666667h-64c-12.8 0-21.333333-8.533333-21.333334-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333334-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM320 874.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM469.333333 874.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333334 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333334 21.333334zM618.666667 874.666667h-64c-12.8 0-21.333333-8.533333-21.333334-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333334-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334z" fill="#E0E0E0" p-id="5116"></path><path d="M768 618.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM768 746.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334zM768 874.666667h-64c-12.8 0-21.333333-8.533333-21.333333-21.333334v-42.666666c0-12.8 8.533333-21.333333 21.333333-21.333334h64c12.8 0 21.333333 8.533333 21.333333 21.333334v42.666666c0 12.8-8.533333 21.333333-21.333333 21.333334z" fill="#BDBDBD" p-id="5117"></path></svg>`)

func init() {
	plugin.AllSystemPlugin = append(plugin.AllSystemPlugin, &CalculatorPlugin{})
}

type CalculatorHistory struct {
	Expression string
	Result     string
	AddDate    string
}

type CalculatorPlugin struct {
	api       plugin.API
	histories []CalculatorHistory
}

func (c *CalculatorPlugin) GetMetadata() plugin.Metadata {
	return plugin.Metadata{
		Id:            "bd723c38-f28d-4152-8621-76fd21d6456e",
		Name:          "Calculator",
		Author:        "Wox Launcher",
		Website:       "https://github.com/Wox-launcher/Wox",
		Version:       "1.0.0",
		MinWoxVersion: "2.0.0",
		Runtime:       "Go",
		Description:   "Calculator for Wox",
		Icon:          calculatorIcon.String(),
		Entry:         "",
		TriggerKeywords: []string{
			"*",
			"calculator",
		},
		Commands: []plugin.MetadataCommand{},
		SupportedOS: []string{
			"Windows",
			"Macos",
			"Linux",
		},
	}
}

func (c *CalculatorPlugin) Init(ctx context.Context, initParams plugin.InitParams) {
	c.api = initParams.API
}

func (c *CalculatorPlugin) Query(ctx context.Context, query plugin.Query) []plugin.QueryResult {
	var results []plugin.QueryResult

	if query.TriggerKeyword == "" {
		//only calculate if query has operators
		if !strings.ContainsAny(query.Search, "+-*/(") {
			return []plugin.QueryResult{}
		}

		val, err := Calculate(query.Search)
		if err != nil {
			return []plugin.QueryResult{}
		}
		result := val.String()

		results = append(results, plugin.QueryResult{
			Title: result,
			Icon:  calculatorIcon,
			Actions: []plugin.QueryResultAction{
				{
					Action: func(ctx context.Context, actionContext plugin.ActionContext) {
						c.histories = append(c.histories, CalculatorHistory{
							Expression: query.Search,
							Result:     result,
							AddDate:    util.FormatDateTime(util.GetSystemTime()),
						})
						clipboard.WriteText(result)
					},
				},
			},
		})
	}

	// only show history if query has trigger keyword
	if query.TriggerKeyword != "" {
		val, err := Calculate(query.Search)
		if err == nil {
			result := val.String()
			results = append(results, plugin.QueryResult{
				Title: result,
				Icon:  calculatorIcon,
				Actions: []plugin.QueryResultAction{
					{
						Action: func(ctx context.Context, actionContext plugin.ActionContext) {
							c.histories = append(c.histories, CalculatorHistory{
								Expression: query.Search,
								Result:     result,
								AddDate:    util.FormatDateTime(util.GetSystemTime()),
							})
							clipboard.WriteText(result)
						},
					},
				},
			})
		}

		//show top 500 histories order by desc
		var count = 0
		for i := len(c.histories) - 1; i >= 0; i-- {
			h := c.histories[i]

			count++
			if count >= 500 {
				break
			}

			if strings.Contains(h.Expression, query.Search) || strings.Contains(h.Result, query.Search) {
				results = append(results, plugin.QueryResult{
					Title:    h.Expression,
					SubTitle: h.Result,
					Icon:     calculatorIcon,
					Actions: []plugin.QueryResultAction{
						{
							Name:      "Copy result",
							IsDefault: true,
							Action: func(ctx context.Context, actionContext plugin.ActionContext) {
								clipboard.WriteText(h.Result)
							},
						},
						{
							Name: "Recalculate",
							Action: func(ctx context.Context, actionContext plugin.ActionContext) {
								c.api.ChangeQuery(ctx, share.PlainQuery{
									QueryType: plugin.QueryTypeInput,
									QueryText: h.Expression,
								})
							},
						},
					},
				})
			}
		}

		if len(results) == 0 {
			results = append(results, plugin.QueryResult{
				Title:   "Input expression to calculate",
				Icon:    calculatorIcon,
				Actions: []plugin.QueryResultAction{},
			})
		}
	}

	return results
}
