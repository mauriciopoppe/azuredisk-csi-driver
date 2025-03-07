/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azuredisk

// about how to get all VM size list,
// refer to https://github.com/kubernetes/kubernetes/issues/77461#issuecomment-492488756
var maxDataDiskCountMap = map[string]int64{
	"BASIC_A0":                  1,
	"BASIC_A1":                  2,
	"BASIC_A2":                  4,
	"BASIC_A3":                  8,
	"BASIC_A4":                  16,
	"STANDARD_A0":               1,
	"STANDARD_A10":              32,
	"STANDARD_A11":              64,
	"STANDARD_A1":               2,
	"STANDARD_A1_V2":            2,
	"STANDARD_A2":               4,
	"STANDARD_A2M_V2":           4,
	"STANDARD_A2_V2":            4,
	"STANDARD_A3":               8,
	"STANDARD_A4":               16,
	"STANDARD_A4M_V2":           8,
	"STANDARD_A4_V2":            8,
	"STANDARD_A5":               4,
	"STANDARD_A6":               8,
	"STANDARD_A7":               16,
	"STANDARD_A8":               32,
	"STANDARD_A8M_V2":           16,
	"STANDARD_A8_V2":            16,
	"STANDARD_A9":               64,
	"STANDARD_B12MS":            16,
	"STANDARD_B16MS":            32,
	"STANDARD_B1LS":             2,
	"STANDARD_B1MS":             2,
	"STANDARD_B1S":              2,
	"STANDARD_B20MS":            32,
	"STANDARD_B2MS":             4,
	"STANDARD_B2S":              4,
	"STANDARD_B4MS":             8,
	"STANDARD_B8MS":             16,
	"STANDARD_D11":              8,
	"STANDARD_D11_V2":           8,
	"STANDARD_D11_V2_PROMO":     8,
	"STANDARD_D12":              16,
	"STANDARD_D12_V2":           16,
	"STANDARD_D12_V2_PROMO":     16,
	"STANDARD_D13":              32,
	"STANDARD_D13_V2":           32,
	"STANDARD_D13_V2_PROMO":     32,
	"STANDARD_D1":               4,
	"STANDARD_D14":              64,
	"STANDARD_D14_V2":           64,
	"STANDARD_D14_V2_PROMO":     64,
	"STANDARD_D15_V2":           64,
	"STANDARD_D16ADS_V5":        32,
	"STANDARD_D16AS_V3":         32,
	"STANDARD_D16AS_V4":         32,
	"STANDARD_D16AS_V5":         32,
	"STANDARD_D16A_V3":          32,
	"STANDARD_D16A_V4":          32,
	"STANDARD_D16DARM_V3":       32,
	"STANDARD_D16DS_V4":         32,
	"STANDARD_D16DS_V5":         32,
	"STANDARD_D16D_V4":          32,
	"STANDARD_D16D_V5":          32,
	"STANDARD_D16PDS_V5":        32,
	"STANDARD_D16PLDS_V5":       32,
	"STANDARD_D16PLS_V5":        32,
	"STANDARD_D16PS_V5":         32,
	"STANDARD_D16S_V3":          32,
	"STANDARD_D16S_V4":          32,
	"STANDARD_D16S_V5":          32,
	"STANDARD_D16_V3":           32,
	"STANDARD_D16_V4":           32,
	"STANDARD_D16_V5":           32,
	"STANDARD_D1_V2":            4,
	"STANDARD_D2":               8,
	"STANDARD_D2ADS_V5":         4,
	"STANDARD_D2AS_V3":          4,
	"STANDARD_D2AS_V4":          4,
	"STANDARD_D2AS_V5":          4,
	"STANDARD_D2A_V3":           4,
	"STANDARD_D2A_V4":           4,
	"STANDARD_D2DARM_V3":        4,
	"STANDARD_D2DS_V4":          4,
	"STANDARD_D2DS_V5":          4,
	"STANDARD_D2D_V4":           4,
	"STANDARD_D2D_V5":           4,
	"STANDARD_D2PDS_V5":         4,
	"STANDARD_D2PLDS_V5":        4,
	"STANDARD_D2PLS_V5":         4,
	"STANDARD_D2PS_V5":          4,
	"STANDARD_D2S_V3":           4,
	"STANDARD_D2S_V4":           4,
	"STANDARD_D2S_V5":           4,
	"STANDARD_D2_V2":            8,
	"STANDARD_D2_V2_PROMO":      8,
	"STANDARD_D2_V3":            4,
	"STANDARD_D2_V4":            4,
	"STANDARD_D2_V5":            4,
	"STANDARD_D3":               16,
	"STANDARD_D32ADS_V5":        32,
	"STANDARD_D32AS_V3":         32,
	"STANDARD_D32AS_V4":         32,
	"STANDARD_D32AS_V5":         32,
	"STANDARD_D32A_V3":          32,
	"STANDARD_D32A_V4":          32,
	"STANDARD_D32DARM_V3":       32,
	"STANDARD_D32DS_V4":         32,
	"STANDARD_D32DS_V5":         32,
	"STANDARD_D32D_V4":          32,
	"STANDARD_D32D_V5":          32,
	"STANDARD_D32PDS_V5":        32,
	"STANDARD_D32PLDS_V5":       32,
	"STANDARD_D32PLS_V5":        32,
	"STANDARD_D32PS_V5":         32,
	"STANDARD_D32S_V3":          32,
	"STANDARD_D32S_V4":          32,
	"STANDARD_D32S_V5":          32,
	"STANDARD_D32_V3":           32,
	"STANDARD_D32_V4":           32,
	"STANDARD_D32_V5":           32,
	"STANDARD_D3_V2":            16,
	"STANDARD_D3_V2_PROMO":      16,
	"STANDARD_D4":               32,
	"STANDARD_D48ADS_V5":        32,
	"STANDARD_D48AS_V3":         32,
	"STANDARD_D48AS_V4":         32,
	"STANDARD_D48AS_V5":         32,
	"STANDARD_D48A_V3":          32,
	"STANDARD_D48A_V4":          32,
	"STANDARD_D48DARM_V3":       32,
	"STANDARD_D48DS_V4":         32,
	"STANDARD_D48DS_V5":         32,
	"STANDARD_D48D_V4":          32,
	"STANDARD_D48D_V5":          32,
	"STANDARD_D48PDS_V5":        32,
	"STANDARD_D48PLDS_V5":       32,
	"STANDARD_D48PLS_V5":        32,
	"STANDARD_D48PS_V5":         32,
	"STANDARD_D48S_V3":          32,
	"STANDARD_D48S_V4":          32,
	"STANDARD_D48S_V5":          32,
	"STANDARD_D48_V3":           32,
	"STANDARD_D48_V4":           32,
	"STANDARD_D48_V5":           32,
	"STANDARD_D4ADS_V5":         8,
	"STANDARD_D4AS_V3":          8,
	"STANDARD_D4AS_V4":          8,
	"STANDARD_D4AS_V5":          8,
	"STANDARD_D4A_V3":           8,
	"STANDARD_D4A_V4":           8,
	"STANDARD_D4DARM_V3":        8,
	"STANDARD_D4DS_V4":          8,
	"STANDARD_D4DS_V5":          8,
	"STANDARD_D4D_V4":           8,
	"STANDARD_D4D_V5":           8,
	"STANDARD_D4PDS_V5":         8,
	"STANDARD_D4PLDS_V5":        8,
	"STANDARD_D4PLS_V5":         8,
	"STANDARD_D4PS_V5":          8,
	"STANDARD_D4S_V3":           8,
	"STANDARD_D4S_V4":           8,
	"STANDARD_D4S_V5":           8,
	"STANDARD_D4_V2":            32,
	"STANDARD_D4_V2_PROMO":      32,
	"STANDARD_D4_V3":            8,
	"STANDARD_D4_V4":            8,
	"STANDARD_D4_V5":            8,
	"STANDARD_D5_V2":            64,
	"STANDARD_D5_V2_PROMO":      64,
	"STANDARD_D64ADS_V5":        32,
	"STANDARD_D64AS_V3":         32,
	"STANDARD_D64AS_V4":         32,
	"STANDARD_D64AS_V5":         32,
	"STANDARD_D64A_V3":          32,
	"STANDARD_D64A_V4":          32,
	"STANDARD_D64DS_V4":         32,
	"STANDARD_D64DS_V5":         32,
	"STANDARD_D64D_V4":          32,
	"STANDARD_D64D_V5":          32,
	"STANDARD_D64PDS_V5":        32,
	"STANDARD_D64PLDS_V5":       32,
	"STANDARD_D64PLS_V5":        32,
	"STANDARD_D64PS_V5":         32,
	"STANDARD_D64S_V3":          32,
	"STANDARD_D64S_V4":          32,
	"STANDARD_D64S_V5":          32,
	"STANDARD_D64_V3":           32,
	"STANDARD_D64_V4":           32,
	"STANDARD_D64_V5":           32,
	"STANDARD_D8ADS_V5":         16,
	"STANDARD_D8AS_V3":          16,
	"STANDARD_D8AS_V4":          16,
	"STANDARD_D8AS_V5":          16,
	"STANDARD_D8A_V3":           16,
	"STANDARD_D8A_V4":           16,
	"STANDARD_D8DARM_V3":        16,
	"STANDARD_D8DS_V4":          16,
	"STANDARD_D8DS_V5":          16,
	"STANDARD_D8D_V4":           16,
	"STANDARD_D8D_V5":           16,
	"STANDARD_D8PDS_V5":         16,
	"STANDARD_D8PLDS_V5":        16,
	"STANDARD_D8PLS_V5":         16,
	"STANDARD_D8PS_V5":          16,
	"STANDARD_D8S_V3":           16,
	"STANDARD_D8S_V4":           16,
	"STANDARD_D8S_V5":           16,
	"STANDARD_D8_V3":            16,
	"STANDARD_D8_V4":            16,
	"STANDARD_D8_V5":            16,
	"STANDARD_D96ADS_V5":        32,
	"STANDARD_D96AS_V4":         32,
	"STANDARD_D96AS_V5":         32,
	"STANDARD_D96A_V4":          32,
	"STANDARD_DC16DMS_V3":       32,
	"STANDARD_DC16MS_V3":        32,
	"STANDARD_DC1DMS_V3":        4,
	"STANDARD_DC1MS_V3":         4,
	"STANDARD_D96DS_V5":         32,
	"STANDARD_D96D_V5":          32,
	"STANDARD_D96S_V5":          32,
	"STANDARD_D96_V5":           32,
	"STANDARD_DC16ADS_V5":       32,
	"STANDARD_DC16AS_V5":        32,
	"STANDARD_DC16DS_V3":        32,
	"STANDARD_DC16S_V3":         32,
	"STANDARD_DC1DS_V3":         4,
	"STANDARD_DC1S_V2":          1,
	"STANDARD_DC24DMS_V3":       32,
	"STANDARD_DC24MS_V3":        32,
	"STANDARD_DC2DMS_V3":        8,
	"STANDARD_DC2MS_V3":         8,
	"STANDARD_DC1S_V3":          4,
	"STANDARD_DC24DS_V3":        32,
	"STANDARD_DC24S_V3":         32,
	"STANDARD_DC2DS_V3":         8,
	"STANDARD_DC2ADS_V5":        4,
	"STANDARD_DC2AS_V5":         4,
	"STANDARD_DC2S":             2,
	"STANDARD_DC2S_V2":          2,
	"STANDARD_DC32DMS_V3":       32,
	"STANDARD_DC32MS_V3":        32,
	"STANDARD_DC4DMS_V3":        16,
	"STANDARD_DC4MS_V3":         16,
	"STANDARD_DC2S_V3":          8,
	"STANDARD_DC32ADS_V5":       32,
	"STANDARD_DC32AS_V5":        32,
	"STANDARD_DC32DS_V3":        32,
	"STANDARD_DC32S_V3":         32,
	"STANDARD_DC48ADS_V5":       32,
	"STANDARD_DC48AS_V5":        32,
	"STANDARD_DC48DS_V3":        32,
	"STANDARD_DC48S_V3":         32,
	"STANDARD_DC4ADS_V5":        8,
	"STANDARD_DC4AS_V5":         8,
	"STANDARD_DC4DS_V3":         16,
	"STANDARD_DC4S":             4,
	"STANDARD_DC4S_V2":          4,
	"STANDARD_DC8DMS_V3":        32,
	"STANDARD_DC8MS_V3":         32,
	"STANDARD_DC8S":             8,
	"STANDARD_DC4S_V3":          16,
	"STANDARD_DC64ADS_V5":       32,
	"STANDARD_DC64AS_V5":        32,
	"STANDARD_DC8ADS_V5":        16,
	"STANDARD_DC8AS_V5":         16,
	"STANDARD_DC8DS_V3":         32,
	"STANDARD_DC8S_V3":          32,
	"STANDARD_DC8_V2":           8,
	"STANDARD_DC96ADS_V5":       32,
	"STANDARD_DC96AS_V5":        32,
	"STANDARD_DS11-1_V2":        8,
	"STANDARD_DS11":             8,
	"STANDARD_DS11_V2":          8,
	"STANDARD_DS11_V2_PROMO":    8,
	"STANDARD_DS12":             16,
	"STANDARD_DS12-1_V2":        16,
	"STANDARD_DS12-2_V2":        16,
	"STANDARD_DS12_V2":          16,
	"STANDARD_DS12_V2_PROMO":    16,
	"STANDARD_DS13-2_V2":        32,
	"STANDARD_DS13":             32,
	"STANDARD_DS13-4_V2":        32,
	"STANDARD_DS13_V2":          32,
	"STANDARD_DS13_V2_PROMO":    32,
	"STANDARD_DS1":              4,
	"STANDARD_DS14-4_V2":        64,
	"STANDARD_DS14":             64,
	"STANDARD_DS14-8_V2":        64,
	"STANDARD_DS14_V2":          64,
	"STANDARD_DS14_V2_PROMO":    64,
	"STANDARD_DS15_V2":          64,
	"STANDARD_DS1_V2":           4,
	"STANDARD_DS2":              8,
	"STANDARD_DS2_V2":           8,
	"STANDARD_DS2_V2_PROMO":     8,
	"STANDARD_DS3":              16,
	"STANDARD_DS3_V2":           16,
	"STANDARD_DS3_V2_PROMO":     16,
	"STANDARD_DS4":              32,
	"STANDARD_DS4_V2":           32,
	"STANDARD_DS4_V2_PROMO":     32,
	"STANDARD_DS5_V2":           64,
	"STANDARD_DS5_V2_PROMO":     64,
	"STANDARD_E104IDS_V5":       64,
	"STANDARD_E104ID_V5":        64,
	"STANDARD_E104IS_V5":        64,
	"STANDARD_E104I_V5":         64,
	"STANDARD_E112IADS_V5":      64,
	"STANDARD_E112IAS_V5":       64,
	"STANDARD_E16-4ADS_V5":      32,
	"STANDARD_E16-4AS_V4":       32,
	"STANDARD_E16-4AS_V5":       32,
	"STANDARD_E16-4DS_V4":       32,
	"STANDARD_E16-4DS_V5":       32,
	"STANDARD_E16-4S_V3":        32,
	"STANDARD_E16-4S_V4":        32,
	"STANDARD_E16-4S_V5":        32,
	"STANDARD_E16-8ADS_V5":      32,
	"STANDARD_E16-8AS_V4":       32,
	"STANDARD_E16-8AS_V5":       32,
	"STANDARD_E16-8DS_V4":       32,
	"STANDARD_E16-8DS_V5":       32,
	"STANDARD_E16-8S_V3":        32,
	"STANDARD_E16-8S_V4":        32,
	"STANDARD_E16-8S_V5":        32,
	"STANDARD_E16ADS_V5":        32,
	"STANDARD_E16AS_V4":         32,
	"STANDARD_E16AS_V5":         32,
	"STANDARD_E16A_V4":          32,
	"STANDARD_E16BDS_V5":        32,
	"STANDARD_E16BS_V5":         32,
	"STANDARD_E16DARM_V3":       32,
	"STANDARD_E16DS_V4":         32,
	"STANDARD_E16DS_V5":         32,
	"STANDARD_E16D_V4":          32,
	"STANDARD_E16D_V5":          32,
	"STANDARD_E16PDS_V5":        32,
	"STANDARD_E16PS_V5":         32,
	"STANDARD_E16S_V3":          32,
	"STANDARD_E16S_V4":          32,
	"STANDARD_E16S_V5":          32,
	"STANDARD_E16_V3":           32,
	"STANDARD_E16_V4":           32,
	"STANDARD_E16_V5":           32,
	"STANDARD_E20ADS_V5":        32,
	"STANDARD_E20AS_V4":         32,
	"STANDARD_E20AS_V5":         32,
	"STANDARD_E20A_V4":          32,
	"STANDARD_E20DARM_V3":       32,
	"STANDARD_E20DS_V4":         32,
	"STANDARD_E20DS_V5":         32,
	"STANDARD_E20D_V4":          32,
	"STANDARD_E20D_V5":          32,
	"STANDARD_E20PDS_V5":        32,
	"STANDARD_E20PS_V5":         32,
	"STANDARD_E20S_V3":          32,
	"STANDARD_E20S_V4":          32,
	"STANDARD_E20S_V5":          32,
	"STANDARD_E20_V3":           32,
	"STANDARD_E20_V4":           32,
	"STANDARD_E20_V5":           32,
	"STANDARD_E2ADS_V5":         4,
	"STANDARD_E2AS_V4":          4,
	"STANDARD_E2AS_V5":          4,
	"STANDARD_E2A_V4":           4,
	"STANDARD_E2BDS_V5":         4,
	"STANDARD_E2BS_V5":          4,
	"STANDARD_E2DARM_V3":        4,
	"STANDARD_E2DS_V4":          4,
	"STANDARD_E2DS_V5":          4,
	"STANDARD_E2D_V4":           4,
	"STANDARD_E2D_V5":           4,
	"STANDARD_E2PDS_V5":         4,
	"STANDARD_E2PS_V5":          4,
	"STANDARD_E2S_V3":           4,
	"STANDARD_E2S_V4":           4,
	"STANDARD_E2S_V5":           4,
	"STANDARD_E2_V3":            4,
	"STANDARD_E2_V4":            4,
	"STANDARD_E2_V5":            4,
	"STANDARD_E32-16ADS_V5":     32,
	"STANDARD_E32-16AS_V4":      32,
	"STANDARD_E32-16AS_V5":      32,
	"STANDARD_E32-16DS_V4":      32,
	"STANDARD_E32-16DS_V5":      32,
	"STANDARD_E32-16S_V3":       32,
	"STANDARD_E32-16S_V4":       32,
	"STANDARD_E32-16S_V5":       32,
	"STANDARD_E32-8ADS_V5":      32,
	"STANDARD_E32-8AS_V4":       32,
	"STANDARD_E32-8AS_V5":       32,
	"STANDARD_E32-8DS_V4":       32,
	"STANDARD_E32-8DS_V5":       32,
	"STANDARD_E32-8S_V3":        32,
	"STANDARD_E32-8S_V4":        32,
	"STANDARD_E32-8S_V5":        32,
	"STANDARD_E32ADS_V5":        32,
	"STANDARD_E32AS_V4":         32,
	"STANDARD_E32AS_V5":         32,
	"STANDARD_E32A_V4":          32,
	"STANDARD_E32BDS_V5":        32,
	"STANDARD_E32BS_V5":         32,
	"STANDARD_E32DARM_V3":       32,
	"STANDARD_E32DS_V4":         32,
	"STANDARD_E32DS_V5":         32,
	"STANDARD_E32D_V4":          32,
	"STANDARD_E32D_V5":          32,
	"STANDARD_E32PDS_V5":        32,
	"STANDARD_E32PS_V5":         32,
	"STANDARD_E32S_V3":          32,
	"STANDARD_E32S_V4":          32,
	"STANDARD_E32S_V5":          32,
	"STANDARD_E32_V3":           32,
	"STANDARD_E32_V4":           32,
	"STANDARD_E32_V5":           32,
	"STANDARD_E4-2ADS_V5":       8,
	"STANDARD_E4-2AS_V4":        8,
	"STANDARD_E4-2AS_V5":        8,
	"STANDARD_E4-2DS_V4":        8,
	"STANDARD_E4-2DS_V5":        8,
	"STANDARD_E4-2S_V3":         8,
	"STANDARD_E4-2S_V4":         8,
	"STANDARD_E4-2S_V5":         8,
	"STANDARD_E48ADS_V5":        32,
	"STANDARD_E48AS_V4":         32,
	"STANDARD_E48AS_V5":         32,
	"STANDARD_E48A_V4":          32,
	"STANDARD_E48BDS_V5":        32,
	"STANDARD_E48BS_V5":         32,
	"STANDARD_E48DARM_V3":       32,
	"STANDARD_E48DS_V4":         32,
	"STANDARD_E48DS_V5":         32,
	"STANDARD_E48D_V4":          32,
	"STANDARD_E48D_V5":          32,
	"STANDARD_E48S_V3":          32,
	"STANDARD_E48S_V4":          32,
	"STANDARD_E48S_V5":          32,
	"STANDARD_E48_V3":           32,
	"STANDARD_E48_V4":           32,
	"STANDARD_E48_V5":           32,
	"STANDARD_E4ADS_V5":         8,
	"STANDARD_E4AS_V4":          8,
	"STANDARD_E4AS_V5":          8,
	"STANDARD_E4A_V4":           8,
	"STANDARD_E4BDS_V5":         8,
	"STANDARD_E4BS_V5":          8,
	"STANDARD_E4DARM_V3":        8,
	"STANDARD_E4DS_V4":          8,
	"STANDARD_E4DS_V5":          8,
	"STANDARD_E4D_V4":           8,
	"STANDARD_E4D_V5":           8,
	"STANDARD_E4PDS_V5":         8,
	"STANDARD_E4PS_V5":          8,
	"STANDARD_E4S_V3":           8,
	"STANDARD_E4S_V4":           8,
	"STANDARD_E4S_V5":           8,
	"STANDARD_E4_V3":            8,
	"STANDARD_E4_V4":            8,
	"STANDARD_E4_V5":            8,
	"STANDARD_E64-16ADS_V5":     32,
	"STANDARD_E64-16AS_V4":      32,
	"STANDARD_E64-16AS_V5":      32,
	"STANDARD_E64-16DS_V4":      32,
	"STANDARD_E64-16DS_V5":      32,
	"STANDARD_E64-16S_V3":       32,
	"STANDARD_E64-16S_V4":       32,
	"STANDARD_E64-16S_V5":       32,
	"STANDARD_E64-32ADS_V5":     32,
	"STANDARD_E64-32AS_V4":      32,
	"STANDARD_E64-32AS_V5":      32,
	"STANDARD_E64-32DS_V4":      32,
	"STANDARD_E64-32DS_V5":      32,
	"STANDARD_E64-32S_V3":       32,
	"STANDARD_E64-32S_V4":       32,
	"STANDARD_E64-32S_V5":       32,
	"STANDARD_E64ADS_V5":        32,
	"STANDARD_E64AS_V4":         32,
	"STANDARD_E64AS_V5":         32,
	"STANDARD_E64A_V4":          32,
	"STANDARD_E64DS_V4":         32,
	"STANDARD_E64DS_V5":         32,
	"STANDARD_E64D_V4":          32,
	"STANDARD_E64D_V5":          32,
	"STANDARD_E64IS_V3":         32,
	"STANDARD_E64I_V3":          32,
	"STANDARD_E64S_V3":          32,
	"STANDARD_E64S_V4":          32,
	"STANDARD_E64S_V5":          32,
	"STANDARD_E64_V3":           32,
	"STANDARD_E64_V4":           32,
	"STANDARD_E64_V5":           32,
	"STANDARD_E80IDS_V4":        64,
	"STANDARD_E80IS_V4":         64,
	"STANDARD_E8-2ADS_V5":       16,
	"STANDARD_E8-2AS_V4":        16,
	"STANDARD_E8-2AS_V5":        16,
	"STANDARD_E8-2DS_V4":        16,
	"STANDARD_E8-2DS_V5":        16,
	"STANDARD_E8-2S_V3":         16,
	"STANDARD_E8-2S_V4":         16,
	"STANDARD_E8-2S_V5":         16,
	"STANDARD_E8-4ADS_V5":       16,
	"STANDARD_E8-4AS_V4":        16,
	"STANDARD_E8-4AS_V5":        16,
	"STANDARD_E8-4DS_V4":        16,
	"STANDARD_E8-4DS_V5":        32,
	"STANDARD_E8-4S_V3":         16,
	"STANDARD_E8-4S_V4":         16,
	"STANDARD_E8-4S_V5":         32,
	"STANDARD_E8ADS_V5":         16,
	"STANDARD_E8AS_V4":          16,
	"STANDARD_E8AS_V5":          16,
	"STANDARD_E8A_V4":           16,
	"STANDARD_E8BDS_V5":         16,
	"STANDARD_E8BS_V5":          16,
	"STANDARD_E8DARM_V3":        16,
	"STANDARD_E8DS_V4":          16,
	"STANDARD_E8DS_V5":          16,
	"STANDARD_E8D_V4":           16,
	"STANDARD_E8D_V5":           16,
	"STANDARD_E8PDS_V5":         16,
	"STANDARD_E8PS_V5":          16,
	"STANDARD_E8S_V3":           16,
	"STANDARD_E8S_V4":           16,
	"STANDARD_E8S_V5":           16,
	"STANDARD_E8_V3":            16,
	"STANDARD_E8_V4":            16,
	"STANDARD_E8_V5":            16,
	"STANDARD_E96-24ADS_V5":     32,
	"STANDARD_E96-24AS_V4":      32,
	"STANDARD_E96-24AS_V5":      32,
	"STANDARD_E96-24DS_V5":      32,
	"STANDARD_E96-24S_V5":       32,
	"STANDARD_E96-48ADS_V5":     32,
	"STANDARD_E96-48AS_V4":      32,
	"STANDARD_E96-48AS_V5":      32,
	"STANDARD_E96-48DS_V5":      32,
	"STANDARD_E96-48S_V5":       32,
	"STANDARD_E96ADS_V5":        32,
	"STANDARD_E96AS_V4":         32,
	"STANDARD_E96AS_V5":         32,
	"STANDARD_E96A_V4":          32,
	"STANDARD_E96DS_V5":         32,
	"STANDARD_E96D_V5":          32,
	"STANDARD_E96IAS_V4":        32,
	"STANDARD_E96S_V5":          32,
	"STANDARD_E96_V5":           32,
	"STANDARD_EC16ADS_V5":       32,
	"STANDARD_EC16AS_V5":        32,
	"STANDARD_EC20ADS_V5":       32,
	"STANDARD_EC20AS_V5":        32,
	"STANDARD_EC2ADS_V5":        4,
	"STANDARD_EC2AS_V5":         4,
	"STANDARD_EC32ADS_V5":       32,
	"STANDARD_EC32AS_V5":        32,
	"STANDARD_EC48ADS_V5":       32,
	"STANDARD_EC48AS_V5":        32,
	"STANDARD_EC4ADS_V5":        8,
	"STANDARD_EC4AS_V5":         8,
	"STANDARD_EC64ADS_V5":       32,
	"STANDARD_EC64AS_V5":        32,
	"STANDARD_EC8ADS_V5":        16,
	"STANDARD_EC8AS_V5":         16,
	"STANDARD_EC96ADS_V5":       32,
	"STANDARD_EC96AS_V5":        32,
	"STANDARD_EC96IADS_V5":      32,
	"STANDARD_EC96IAS_V5":       32,
	"STANDARD_F1":               4,
	"STANDARD_F16":              64,
	"STANDARD_F16S":             64,
	"STANDARD_F16S_V2":          32,
	"STANDARD_F1S":              4,
	"STANDARD_F2":               8,
	"STANDARD_F2S":              8,
	"STANDARD_F2S_V2":           4,
	"STANDARD_F32S_V2":          32,
	"STANDARD_F4":               16,
	"STANDARD_F48S_V2":          32,
	"STANDARD_F4S":              16,
	"STANDARD_F4S_V2":           8,
	"STANDARD_F64S_V2":          32,
	"STANDARD_F72S_V2":          32,
	"STANDARD_F8":               32,
	"STANDARD_F8S":              32,
	"STANDARD_F8S_V2":           16,
	"STANDARD_FX12MDS":          24,
	"STANDARD_FX24MDS":          32,
	"STANDARD_FX36MDS":          32,
	"STANDARD_FX48MDS":          32,
	"STANDARD_FX4MDS":           8,
	"STANDARD_G1":               8,
	"STANDARD_G2":               16,
	"STANDARD_G3":               32,
	"STANDARD_G4":               64,
	"STANDARD_G5":               64,
	"STANDARD_GS1":              8,
	"STANDARD_GS2":              16,
	"STANDARD_GS3":              32,
	"STANDARD_GS4-4":            64,
	"STANDARD_GS4":              64,
	"STANDARD_GS4-8":            64,
	"STANDARD_GS5-16":           64,
	"STANDARD_GS5":              64,
	"STANDARD_GS5-8":            64,
	"STANDARD_H16":              64,
	"STANDARD_H16M":             64,
	"STANDARD_H16M_PROMO":       64,
	"STANDARD_H16MR":            64,
	"STANDARD_H16MR_PROMO":      64,
	"STANDARD_H16_PROMO":        64,
	"STANDARD_H16R":             64,
	"STANDARD_H16R_PROMO":       64,
	"STANDARD_H8":               32,
	"STANDARD_H8M":              32,
	"STANDARD_H8M_PROMO":        32,
	"STANDARD_H8_PROMO":         32,
	"STANDARD_HB120-16RS_V2":    32,
	"STANDARD_HB120-16RS_V3":    32,
	"STANDARD_HB120-32RS_V2":    32,
	"STANDARD_HB120-32RS_V3":    32,
	"STANDARD_HB120-64RS_V2":    32,
	"STANDARD_HB120-64RS_V3":    32,
	"STANDARD_HB120-96RS_V2":    32,
	"STANDARD_HB120-96RS_V3":    32,
	"STANDARD_HB120RS_V2":       8,
	"STANDARD_HB120RS_V3":       32,
	"STANDARD_HB60-15RS":        4,
	"STANDARD_HB60-30RS":        4,
	"STANDARD_HB60-45RS":        4,
	"STANDARD_HB60RS":           4,
	"STANDARD_HC44-16RS":        4,
	"STANDARD_HC44-32RS":        4,
	"STANDARD_HC44RS":           4,
	"STANDARD_L16AS_V3":         32,
	"STANDARD_L16S":             64,
	"STANDARD_L16S_V2":          32,
	"STANDARD_L16S_V3":          32,
	"STANDARD_L32AS_V3":         32,
	"STANDARD_L32S":             64,
	"STANDARD_L32S_V2":          32,
	"STANDARD_L32S_V3":          32,
	"STANDARD_L48AS_V3":         32,
	"STANDARD_L48S_V2":          32,
	"STANDARD_L48S_V3":          32,
	"STANDARD_L4S":              16,
	"STANDARD_L64AS_V3":         32,
	"STANDARD_L64S_V2":          32,
	"STANDARD_L64S_V3":          32,
	"STANDARD_L80AS_V3":         32,
	"STANDARD_L80S_V2":          32,
	"STANDARD_L80S_V3":          32,
	"STANDARD_L8AS_V3":          16,
	"STANDARD_L8S":              32,
	"STANDARD_L8S_V2":           16,
	"STANDARD_L8S_V3":           16,
	"STANDARD_M128-32MS":        64,
	"STANDARD_M128":             64,
	"STANDARD_M128-64MS":        64,
	"STANDARD_M128DMS_V2":       64,
	"STANDARD_M128DS_V2":        64,
	"STANDARD_M128M":            64,
	"STANDARD_M128MS":           64,
	"STANDARD_M128MS_V2":        64,
	"STANDARD_M128S":            64,
	"STANDARD_M128S_V2":         64,
	"STANDARD_M16-4MS":          16,
	"STANDARD_M16-8MS":          16,
	"STANDARD_M16MS":            16,
	"STANDARD_M192IDMS_V2":      64,
	"STANDARD_M192IDS_V2":       64,
	"STANDARD_M192IMS_V2":       64,
	"STANDARD_M192IS_V2":        64,
	"STANDARD_M208MS_V2":        64,
	"STANDARD_M208S_V2":         64,
	"STANDARD_M32-16MS":         32,
	"STANDARD_M32-8MS":          32,
	"STANDARD_M32DMS_V2":        32,
	"STANDARD_M32LS":            32,
	"STANDARD_M32MS":            32,
	"STANDARD_M32MS_V2":         32,
	"STANDARD_M32TS":            32,
	"STANDARD_M416-208MS_V2":    64,
	"STANDARD_M416-208S_V2":     64,
	"STANDARD_M416MS_V2":        64,
	"STANDARD_M416S_V2":         64,
	"STANDARD_M64-16MS":         64,
	"STANDARD_M64-32MS":         64,
	"STANDARD_M64":              64,
	"STANDARD_M64DMS_V2":        64,
	"STANDARD_M64DS_V2":         64,
	"STANDARD_M64LS":            64,
	"STANDARD_M64M":             64,
	"STANDARD_M64MS":            64,
	"STANDARD_M64MS_V2":         64,
	"STANDARD_M64S":             64,
	"STANDARD_M64S_V2":          64,
	"STANDARD_M8-2MS":           8,
	"STANDARD_M832IXS":          64,
	"STANDARD_M8-4MS":           8,
	"STANDARD_M864IXS":          64,
	"STANDARD_M8MS":             8,
	"STANDARD_NC12":             48,
	"STANDARD_NC12_PROMO":       48,
	"STANDARD_NC12S_V2":         24,
	"STANDARD_NC12S_V3":         24,
	"STANDARD_NC16AS_T4_V3":     32,
	"STANDARD_NC24":             64,
	"STANDARD_NC24ADS_A100_V4":  8,
	"STANDARD_NC24_PROMO":       64,
	"STANDARD_NC24R":            64,
	"STANDARD_NC24R_PROMO":      64,
	"STANDARD_NC24RS_V2":        32,
	"STANDARD_NC24RS_V3":        32,
	"STANDARD_NC24S_V2":         32,
	"STANDARD_NC24S_V3":         32,
	"STANDARD_NC32ADS_A10_V4":   32,
	"STANDARD_NC48ADS_A100_V4":  16,
	"STANDARD_NC4AS_T4_V3":      8,
	"STANDARD_NC6":              24,
	"STANDARD_NC64AS_T4_V3":     32,
	"STANDARD_NC6_PROMO":        24,
	"STANDARD_NC6S_V2":          12,
	"STANDARD_NC6S_V3":          12,
	"STANDARD_NC8ADS_A10_V4":    8,
	"STANDARD_NC8AS_T4_V3":      16,
	"STANDARD_NC96ADS_A100_V4":  32,
	"STANDARD_NCC24ADS_A100_V4": 8,
	"STANDARD_ND12S":            24,
	"STANDARD_ND24RS":           32,
	"STANDARD_ND24S":            32,
	"STANDARD_ND40RS_V2":        8,
	"STANDARD_ND40S_V3":         32,
	"STANDARD_ND6S":             12,
	"STANDARD_ND96AMSR_A100_V4": 16,
	"STANDARD_ND96ASR_V4":       16,
	"STANDARD_NP10S":            8,
	"STANDARD_NP20S":            16,
	"STANDARD_NP40S":            32,
	"STANDARD_NV12":             48,
	"STANDARD_NV12ADS_A10_V5":   8,
	"STANDARD_NV12_PROMO":       48,
	"STANDARD_NV12S_V2":         24,
	"STANDARD_NV12S_V3":         12,
	"STANDARD_NV16AS_V4":        32,
	"STANDARD_NV18ADS_A10_V5":   16,
	"STANDARD_NV24":             64,
	"STANDARD_NV24_PROMO":       64,
	"STANDARD_NV24S_V2":         32,
	"STANDARD_NV24S_V3":         24,
	"STANDARD_NV32AS_V4":        32,
	"STANDARD_NV36ADMS_A10_V5":  32,
	"STANDARD_NV36ADS_A10_V5":   32,
	"STANDARD_NV48S_V3":         32,
	"STANDARD_NV4AS_V4":         8,
	"STANDARD_NV6":              24,
	"STANDARD_NV6ADS_A10_V5":    4,
	"STANDARD_NV6_PROMO":        24,
	"STANDARD_NV6S_V2":          12,
	"STANDARD_NV72ADS_A10_V5":   32,
	"STANDARD_NV8AS_V4":         16,
	"STANDARD_PB6S":             12,
}
