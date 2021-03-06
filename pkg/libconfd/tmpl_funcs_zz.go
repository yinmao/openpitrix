// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Auto generated by go generate, DO NOT EDIT!!!

package libconfd

import (
	"text/template"
)

func init() {
	_TemplateFunc_initFuncMap = func(p *TemplateFunc) {
		p.FuncMap = template.FuncMap{
			"add":            p.Add,
			"atoi":           p.Atoi,
			"base":           p.Base,
			"base64Decode":   p.Base64Decode,
			"base64Encode":   p.Base64Encode,
			"cget":           p.Cget,
			"cgets":          p.Cgets,
			"cgetv":          p.Cgetv,
			"cgetvs":         p.Cgetvs,
			"contains":       p.Contains,
			"datetime":       p.Datetime,
			"dir":            p.Dir,
			"div":            p.Div,
			"exists":         p.Exists,
			"fileExists":     p.FileExists,
			"get":            p.Get,
			"getenv":         p.Getenv,
			"gets":           p.Gets,
			"getv":           p.Getv,
			"getvs":          p.Getvs,
			"join":           p.Join,
			"json":           p.Json,
			"jsonArray":      p.JsonArray,
			"lookupIP":       p.LookupIP,
			"lookupIPV4":     p.LookupIPV4,
			"lookupIPV6":     p.LookupIPV6,
			"lookupSRV":      p.LookupSRV,
			"ls":             p.Ls,
			"lsdir":          p.Lsdir,
			"map":            p.Map,
			"mod":            p.Mod,
			"mul":            p.Mul,
			"parseBool":      p.ParseBool,
			"replace":        p.Replace,
			"reverse":        p.Reverse,
			"seq":            p.Seq,
			"sortByLength":   p.SortByLength,
			"sortKVByLength": p.SortKVByLength,
			"split":          p.Split,
			"sub":            p.Sub,
			"toLower":        p.ToLower,
			"toUpper":        p.ToUpper,
			"trimSuffix":     p.TrimSuffix,
		}
	}
}
