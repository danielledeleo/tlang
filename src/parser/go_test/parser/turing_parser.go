// Code generated from Turing.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Turing

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 108, 590,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 3, 2, 6, 2, 124, 10, 2, 13, 2, 14, 2, 125, 3, 3, 3, 3,
	5, 3, 130, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 136, 10, 4, 3, 4, 5, 4,
	139, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 145, 10, 5, 3, 5, 5, 5, 148,
	10, 5, 3, 5, 5, 5, 151, 10, 5, 3, 5, 3, 5, 3, 5, 3, 6, 5, 6, 157, 10, 6,
	3, 6, 3, 6, 5, 6, 161, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 6, 9, 173, 10, 9, 13, 9, 14, 9, 174, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 184, 10, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 193, 10, 11, 3, 12, 5, 12, 196, 10, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 206, 10,
	13, 12, 13, 14, 13, 209, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 218, 10, 14, 3, 15, 3, 15, 5, 15, 222, 10, 15, 3, 15,
	3, 15, 3, 15, 5, 15, 227, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 233,
	10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 239, 10, 16, 3, 17, 3, 17, 3,
	18, 5, 18, 244, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 251,
	10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 259, 10, 19, 12,
	19, 14, 19, 262, 11, 19, 3, 20, 6, 20, 265, 10, 20, 13, 20, 14, 20, 266,
	3, 21, 3, 21, 5, 21, 271, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 5, 22, 279, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23,
	287, 10, 23, 12, 23, 14, 23, 290, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 303, 10, 26, 3, 26,
	3, 26, 3, 26, 5, 26, 308, 10, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 7, 26, 317, 10, 26, 12, 26, 14, 26, 320, 11, 26, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 329, 10, 27, 3, 28, 3, 28, 3,
	28, 5, 28, 334, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 340, 10, 28,
	7, 28, 342, 10, 28, 12, 28, 14, 28, 345, 11, 28, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 355, 10, 30, 12, 30, 14, 30, 358,
	11, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 373, 10, 31, 3, 31, 7, 31, 376, 10,
	31, 12, 31, 14, 31, 379, 11, 31, 3, 32, 3, 32, 3, 32, 5, 32, 384, 10, 32,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 392, 10, 33, 12, 33, 14,
	33, 395, 11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 403,
	10, 34, 12, 34, 14, 34, 406, 11, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 7, 35, 414, 10, 35, 12, 35, 14, 35, 417, 11, 35, 3, 36, 3, 36, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 427, 10, 37, 12, 37, 14,
	37, 430, 11, 37, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 436, 10, 38, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 448,
	10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 7, 43, 465, 10, 43, 12, 43, 14,
	43, 468, 11, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46,
	3, 46, 3, 46, 5, 46, 480, 10, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 5,
	47, 487, 10, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 50, 3, 50, 5, 50, 511, 10, 50, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 7, 52, 523, 10, 52, 12, 52, 14,
	52, 526, 11, 52, 3, 53, 3, 53, 3, 53, 3, 53, 5, 53, 532, 10, 53, 3, 54,
	3, 54, 6, 54, 536, 10, 54, 13, 54, 14, 54, 537, 3, 54, 3, 54, 3, 54, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56,
	553, 10, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 560, 10, 56, 3,
	57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	5, 58, 573, 10, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 7, 59, 581,
	10, 59, 12, 59, 14, 59, 584, 11, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61,
	2, 16, 24, 36, 44, 50, 54, 58, 60, 64, 66, 68, 72, 84, 102, 116, 62, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110,
	112, 114, 116, 118, 120, 2, 10, 4, 2, 13, 13, 16, 16, 3, 2, 58, 60, 4,
	2, 3, 5, 12, 12, 4, 2, 74, 78, 90, 91, 4, 2, 72, 73, 92, 92, 5, 2, 71,
	71, 93, 97, 102, 103, 3, 2, 6, 7, 3, 2, 107, 108, 2, 617, 2, 123, 3, 2,
	2, 2, 4, 129, 3, 2, 2, 2, 6, 131, 3, 2, 2, 2, 8, 140, 3, 2, 2, 2, 10, 156,
	3, 2, 2, 2, 12, 162, 3, 2, 2, 2, 14, 167, 3, 2, 2, 2, 16, 169, 3, 2, 2,
	2, 18, 183, 3, 2, 2, 2, 20, 192, 3, 2, 2, 2, 22, 195, 3, 2, 2, 2, 24, 199,
	3, 2, 2, 2, 26, 217, 3, 2, 2, 2, 28, 232, 3, 2, 2, 2, 30, 238, 3, 2, 2,
	2, 32, 240, 3, 2, 2, 2, 34, 250, 3, 2, 2, 2, 36, 252, 3, 2, 2, 2, 38, 264,
	3, 2, 2, 2, 40, 270, 3, 2, 2, 2, 42, 278, 3, 2, 2, 2, 44, 280, 3, 2, 2,
	2, 46, 291, 3, 2, 2, 2, 48, 295, 3, 2, 2, 2, 50, 302, 3, 2, 2, 2, 52, 328,
	3, 2, 2, 2, 54, 333, 3, 2, 2, 2, 56, 346, 3, 2, 2, 2, 58, 348, 3, 2, 2,
	2, 60, 359, 3, 2, 2, 2, 62, 383, 3, 2, 2, 2, 64, 385, 3, 2, 2, 2, 66, 396,
	3, 2, 2, 2, 68, 407, 3, 2, 2, 2, 70, 418, 3, 2, 2, 2, 72, 420, 3, 2, 2,
	2, 74, 435, 3, 2, 2, 2, 76, 447, 3, 2, 2, 2, 78, 449, 3, 2, 2, 2, 80, 453,
	3, 2, 2, 2, 82, 456, 3, 2, 2, 2, 84, 458, 3, 2, 2, 2, 86, 469, 3, 2, 2,
	2, 88, 473, 3, 2, 2, 2, 90, 476, 3, 2, 2, 2, 92, 483, 3, 2, 2, 2, 94, 490,
	3, 2, 2, 2, 96, 493, 3, 2, 2, 2, 98, 510, 3, 2, 2, 2, 100, 512, 3, 2, 2,
	2, 102, 516, 3, 2, 2, 2, 104, 527, 3, 2, 2, 2, 106, 533, 3, 2, 2, 2, 108,
	542, 3, 2, 2, 2, 110, 559, 3, 2, 2, 2, 112, 561, 3, 2, 2, 2, 114, 572,
	3, 2, 2, 2, 116, 574, 3, 2, 2, 2, 118, 585, 3, 2, 2, 2, 120, 587, 3, 2,
	2, 2, 122, 124, 5, 4, 3, 2, 123, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2,
	125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 3, 3, 2, 2, 2, 127, 130,
	5, 40, 21, 2, 128, 130, 5, 16, 9, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3,
	2, 2, 2, 130, 5, 3, 2, 2, 2, 131, 132, 9, 2, 2, 2, 132, 138, 7, 105, 2,
	2, 133, 135, 7, 64, 2, 2, 134, 136, 5, 36, 19, 2, 135, 134, 3, 2, 2, 2,
	135, 136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 139, 7, 65, 2, 2, 138,
	133, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 7, 3, 2, 2, 2, 140, 141, 7,
	14, 2, 2, 141, 147, 7, 105, 2, 2, 142, 144, 7, 64, 2, 2, 143, 145, 5, 36,
	19, 2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2,
	146, 148, 7, 65, 2, 2, 147, 142, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148,
	150, 3, 2, 2, 2, 149, 151, 7, 105, 2, 2, 150, 149, 3, 2, 2, 2, 150, 151,
	3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 7, 63, 2, 2, 153, 154, 5, 98,
	50, 2, 154, 9, 3, 2, 2, 2, 155, 157, 5, 14, 8, 2, 156, 155, 3, 2, 2, 2,
	156, 157, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 161, 5, 8, 5, 2, 159,
	161, 5, 6, 4, 2, 160, 158, 3, 2, 2, 2, 160, 159, 3, 2, 2, 2, 161, 11, 3,
	2, 2, 2, 162, 163, 5, 10, 6, 2, 163, 164, 5, 38, 20, 2, 164, 165, 7, 8,
	2, 2, 165, 166, 7, 105, 2, 2, 166, 13, 3, 2, 2, 2, 167, 168, 9, 3, 2, 2,
	168, 15, 3, 2, 2, 2, 169, 170, 7, 15, 2, 2, 170, 172, 7, 105, 2, 2, 171,
	173, 5, 18, 10, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 172,
	3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 177, 7, 8,
	2, 2, 177, 178, 7, 105, 2, 2, 178, 17, 3, 2, 2, 2, 179, 184, 5, 20, 11,
	2, 180, 184, 5, 26, 14, 2, 181, 184, 5, 28, 15, 2, 182, 184, 5, 40, 21,
	2, 183, 179, 3, 2, 2, 2, 183, 180, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183,
	182, 3, 2, 2, 2, 184, 19, 3, 2, 2, 2, 185, 186, 7, 38, 2, 2, 186, 193,
	5, 24, 13, 2, 187, 188, 7, 38, 2, 2, 188, 189, 7, 64, 2, 2, 189, 190, 5,
	24, 13, 2, 190, 191, 7, 65, 2, 2, 191, 193, 3, 2, 2, 2, 192, 185, 3, 2,
	2, 2, 192, 187, 3, 2, 2, 2, 193, 21, 3, 2, 2, 2, 194, 196, 5, 32, 17, 2,
	195, 194, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197,
	198, 7, 105, 2, 2, 198, 23, 3, 2, 2, 2, 199, 200, 8, 13, 1, 2, 200, 201,
	5, 22, 12, 2, 201, 207, 3, 2, 2, 2, 202, 203, 12, 3, 2, 2, 203, 204, 7,
	68, 2, 2, 204, 206, 5, 22, 12, 2, 205, 202, 3, 2, 2, 2, 206, 209, 3, 2,
	2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 25, 3, 2, 2, 2,
	209, 207, 3, 2, 2, 2, 210, 211, 7, 40, 2, 2, 211, 218, 5, 30, 16, 2, 212,
	213, 7, 40, 2, 2, 213, 214, 7, 64, 2, 2, 214, 215, 5, 30, 16, 2, 215, 216,
	7, 65, 2, 2, 216, 218, 3, 2, 2, 2, 217, 210, 3, 2, 2, 2, 217, 212, 3, 2,
	2, 2, 218, 27, 3, 2, 2, 2, 219, 221, 7, 41, 2, 2, 220, 222, 7, 42, 2, 2,
	221, 220, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223,
	233, 5, 30, 16, 2, 224, 226, 7, 41, 2, 2, 225, 227, 7, 42, 2, 2, 226, 225,
	3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 229, 7, 64,
	2, 2, 229, 230, 5, 30, 16, 2, 230, 231, 7, 65, 2, 2, 231, 233, 3, 2, 2,
	2, 232, 219, 3, 2, 2, 2, 232, 224, 3, 2, 2, 2, 233, 29, 3, 2, 2, 2, 234,
	239, 7, 105, 2, 2, 235, 236, 7, 105, 2, 2, 236, 237, 7, 89, 2, 2, 237,
	239, 7, 6, 2, 2, 238, 234, 3, 2, 2, 2, 238, 235, 3, 2, 2, 2, 239, 31, 3,
	2, 2, 2, 240, 241, 9, 4, 2, 2, 241, 33, 3, 2, 2, 2, 242, 244, 7, 12, 2,
	2, 243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245,
	246, 5, 116, 59, 2, 246, 247, 7, 63, 2, 2, 247, 248, 5, 98, 50, 2, 248,
	251, 3, 2, 2, 2, 249, 251, 5, 10, 6, 2, 250, 243, 3, 2, 2, 2, 250, 249,
	3, 2, 2, 2, 251, 35, 3, 2, 2, 2, 252, 253, 8, 19, 1, 2, 253, 254, 5, 34,
	18, 2, 254, 260, 3, 2, 2, 2, 255, 256, 12, 3, 2, 2, 256, 257, 7, 68, 2,
	2, 257, 259, 5, 34, 18, 2, 258, 255, 3, 2, 2, 2, 259, 262, 3, 2, 2, 2,
	260, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 37, 3, 2, 2, 2, 262, 260,
	3, 2, 2, 2, 263, 265, 5, 40, 21, 2, 264, 263, 3, 2, 2, 2, 265, 266, 3,
	2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 39, 3, 2, 2,
	2, 268, 271, 5, 76, 39, 2, 269, 271, 5, 74, 38, 2, 270, 268, 3, 2, 2, 2,
	270, 269, 3, 2, 2, 2, 271, 41, 3, 2, 2, 2, 272, 279, 5, 118, 60, 2, 273,
	279, 7, 105, 2, 2, 274, 275, 7, 64, 2, 2, 275, 276, 5, 70, 36, 2, 276,
	277, 7, 65, 2, 2, 277, 279, 3, 2, 2, 2, 278, 272, 3, 2, 2, 2, 278, 273,
	3, 2, 2, 2, 278, 274, 3, 2, 2, 2, 279, 43, 3, 2, 2, 2, 280, 281, 8, 23,
	1, 2, 281, 282, 5, 70, 36, 2, 282, 288, 3, 2, 2, 2, 283, 284, 12, 3, 2,
	2, 284, 285, 7, 68, 2, 2, 285, 287, 5, 44, 23, 4, 286, 283, 3, 2, 2, 2,
	287, 290, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289,
	45, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 291, 292, 5, 42, 22, 2, 292, 293,
	7, 79, 2, 2, 293, 294, 5, 42, 22, 2, 294, 47, 3, 2, 2, 2, 295, 296, 7,
	62, 2, 2, 296, 297, 5, 42, 22, 2, 297, 49, 3, 2, 2, 2, 298, 299, 8, 26,
	1, 2, 299, 303, 5, 42, 22, 2, 300, 303, 5, 48, 25, 2, 301, 303, 5, 46,
	24, 2, 302, 298, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 301, 3, 2, 2, 2,
	303, 318, 3, 2, 2, 2, 304, 305, 12, 5, 2, 2, 305, 307, 7, 64, 2, 2, 306,
	308, 5, 44, 23, 2, 307, 306, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309,
	3, 2, 2, 2, 309, 317, 7, 65, 2, 2, 310, 311, 12, 4, 2, 2, 311, 312, 7,
	66, 2, 2, 312, 317, 7, 105, 2, 2, 313, 314, 12, 3, 2, 2, 314, 315, 7, 70,
	2, 2, 315, 317, 7, 105, 2, 2, 316, 304, 3, 2, 2, 2, 316, 310, 3, 2, 2,
	2, 316, 313, 3, 2, 2, 2, 317, 320, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 318,
	319, 3, 2, 2, 2, 319, 51, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 321, 329, 5,
	50, 26, 2, 322, 323, 7, 72, 2, 2, 323, 329, 5, 52, 27, 2, 324, 325, 7,
	73, 2, 2, 325, 329, 5, 52, 27, 2, 326, 327, 7, 69, 2, 2, 327, 329, 5, 52,
	27, 2, 328, 321, 3, 2, 2, 2, 328, 322, 3, 2, 2, 2, 328, 324, 3, 2, 2, 2,
	328, 326, 3, 2, 2, 2, 329, 53, 3, 2, 2, 2, 330, 331, 8, 28, 1, 2, 331,
	334, 5, 50, 26, 2, 332, 334, 5, 52, 27, 2, 333, 330, 3, 2, 2, 2, 333, 332,
	3, 2, 2, 2, 334, 343, 3, 2, 2, 2, 335, 336, 12, 3, 2, 2, 336, 339, 5, 56,
	29, 2, 337, 340, 5, 50, 26, 2, 338, 340, 5, 52, 27, 2, 339, 337, 3, 2,
	2, 2, 339, 338, 3, 2, 2, 2, 340, 342, 3, 2, 2, 2, 341, 335, 3, 2, 2, 2,
	342, 345, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344,
	55, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 346, 347, 9, 5, 2, 2, 347, 57, 3,
	2, 2, 2, 348, 349, 8, 30, 1, 2, 349, 350, 5, 54, 28, 2, 350, 356, 3, 2,
	2, 2, 351, 352, 12, 3, 2, 2, 352, 353, 9, 6, 2, 2, 353, 355, 5, 54, 28,
	2, 354, 351, 3, 2, 2, 2, 355, 358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356,
	357, 3, 2, 2, 2, 357, 59, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 360, 8,
	31, 1, 2, 360, 361, 5, 58, 30, 2, 361, 377, 3, 2, 2, 2, 362, 372, 12, 3,
	2, 2, 363, 373, 7, 80, 2, 2, 364, 373, 7, 81, 2, 2, 365, 373, 7, 82, 2,
	2, 366, 373, 7, 83, 2, 2, 367, 373, 7, 84, 2, 2, 368, 373, 7, 85, 2, 2,
	369, 373, 7, 89, 2, 2, 370, 371, 7, 61, 2, 2, 371, 373, 7, 89, 2, 2, 372,
	363, 3, 2, 2, 2, 372, 364, 3, 2, 2, 2, 372, 365, 3, 2, 2, 2, 372, 366,
	3, 2, 2, 2, 372, 367, 3, 2, 2, 2, 372, 368, 3, 2, 2, 2, 372, 369, 3, 2,
	2, 2, 372, 370, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 376, 5, 58, 30,
	2, 375, 362, 3, 2, 2, 2, 376, 379, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377,
	378, 3, 2, 2, 2, 378, 61, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 380, 384, 5,
	60, 31, 2, 381, 382, 7, 61, 2, 2, 382, 384, 5, 62, 32, 2, 383, 380, 3,
	2, 2, 2, 383, 381, 3, 2, 2, 2, 384, 63, 3, 2, 2, 2, 385, 386, 8, 33, 1,
	2, 386, 387, 5, 62, 32, 2, 387, 393, 3, 2, 2, 2, 388, 389, 12, 3, 2, 2,
	389, 390, 7, 86, 2, 2, 390, 392, 5, 62, 32, 2, 391, 388, 3, 2, 2, 2, 392,
	395, 3, 2, 2, 2, 393, 391, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 65, 3,
	2, 2, 2, 395, 393, 3, 2, 2, 2, 396, 397, 8, 34, 1, 2, 397, 398, 5, 64,
	33, 2, 398, 404, 3, 2, 2, 2, 399, 400, 12, 3, 2, 2, 400, 401, 7, 87, 2,
	2, 401, 403, 5, 64, 33, 2, 402, 399, 3, 2, 2, 2, 403, 406, 3, 2, 2, 2,
	404, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 67, 3, 2, 2, 2, 406, 404,
	3, 2, 2, 2, 407, 408, 8, 35, 1, 2, 408, 409, 5, 66, 34, 2, 409, 415, 3,
	2, 2, 2, 410, 411, 12, 3, 2, 2, 411, 412, 7, 88, 2, 2, 412, 414, 5, 66,
	34, 2, 413, 410, 3, 2, 2, 2, 414, 417, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2,
	415, 416, 3, 2, 2, 2, 416, 69, 3, 2, 2, 2, 417, 415, 3, 2, 2, 2, 418, 419,
	5, 68, 35, 2, 419, 71, 3, 2, 2, 2, 420, 421, 8, 37, 1, 2, 421, 422, 5,
	70, 36, 2, 422, 428, 3, 2, 2, 2, 423, 424, 12, 3, 2, 2, 424, 425, 7, 68,
	2, 2, 425, 427, 5, 70, 36, 2, 426, 423, 3, 2, 2, 2, 427, 430, 3, 2, 2,
	2, 428, 426, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 73, 3, 2, 2, 2, 430,
	428, 3, 2, 2, 2, 431, 436, 5, 96, 49, 2, 432, 436, 5, 110, 56, 2, 433,
	436, 5, 114, 58, 2, 434, 436, 5, 12, 7, 2, 435, 431, 3, 2, 2, 2, 435, 432,
	3, 2, 2, 2, 435, 433, 3, 2, 2, 2, 435, 434, 3, 2, 2, 2, 436, 75, 3, 2,
	2, 2, 437, 448, 5, 70, 36, 2, 438, 448, 5, 78, 40, 2, 439, 448, 5, 80,
	41, 2, 440, 448, 7, 19, 2, 2, 441, 448, 5, 86, 44, 2, 442, 448, 7, 26,
	2, 2, 443, 448, 5, 88, 45, 2, 444, 448, 5, 90, 46, 2, 445, 448, 5, 92,
	47, 2, 446, 448, 5, 94, 48, 2, 447, 437, 3, 2, 2, 2, 447, 438, 3, 2, 2,
	2, 447, 439, 3, 2, 2, 2, 447, 440, 3, 2, 2, 2, 447, 441, 3, 2, 2, 2, 447,
	442, 3, 2, 2, 2, 447, 443, 3, 2, 2, 2, 447, 444, 3, 2, 2, 2, 447, 445,
	3, 2, 2, 2, 447, 446, 3, 2, 2, 2, 448, 77, 3, 2, 2, 2, 449, 450, 5, 50,
	26, 2, 450, 451, 9, 7, 2, 2, 451, 452, 5, 70, 36, 2, 452, 79, 3, 2, 2,
	2, 453, 454, 7, 43, 2, 2, 454, 455, 5, 84, 43, 2, 455, 81, 3, 2, 2, 2,
	456, 457, 5, 76, 39, 2, 457, 83, 3, 2, 2, 2, 458, 459, 8, 43, 1, 2, 459,
	460, 5, 82, 42, 2, 460, 466, 3, 2, 2, 2, 461, 462, 12, 3, 2, 2, 462, 463,
	7, 68, 2, 2, 463, 465, 5, 82, 42, 2, 464, 461, 3, 2, 2, 2, 465, 468, 3,
	2, 2, 2, 466, 464, 3, 2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 85, 3, 2, 2,
	2, 468, 466, 3, 2, 2, 2, 469, 470, 7, 25, 2, 2, 470, 471, 5, 40, 21, 2,
	471, 472, 7, 8, 2, 2, 472, 87, 3, 2, 2, 2, 473, 474, 7, 27, 2, 2, 474,
	475, 5, 70, 36, 2, 475, 89, 3, 2, 2, 2, 476, 479, 7, 28, 2, 2, 477, 478,
	7, 105, 2, 2, 478, 480, 7, 68, 2, 2, 479, 477, 3, 2, 2, 2, 479, 480, 3,
	2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 482, 7, 105, 2, 2, 482, 91, 3, 2, 2,
	2, 483, 486, 7, 29, 2, 2, 484, 485, 7, 105, 2, 2, 485, 487, 7, 68, 2, 2,
	486, 484, 3, 2, 2, 2, 486, 487, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488,
	489, 7, 105, 2, 2, 489, 93, 3, 2, 2, 2, 490, 491, 7, 31, 2, 2, 491, 492,
	5, 70, 36, 2, 492, 95, 3, 2, 2, 2, 493, 494, 7, 11, 2, 2, 494, 495, 7,
	105, 2, 2, 495, 496, 7, 63, 2, 2, 496, 497, 5, 98, 50, 2, 497, 97, 3, 2,
	2, 2, 498, 511, 7, 44, 2, 2, 499, 511, 7, 45, 2, 2, 500, 511, 7, 46, 2,
	2, 501, 511, 7, 53, 2, 2, 502, 511, 7, 54, 2, 2, 503, 511, 7, 55, 2, 2,
	504, 511, 7, 56, 2, 2, 505, 511, 7, 57, 2, 2, 506, 511, 5, 100, 51, 2,
	507, 511, 5, 104, 53, 2, 508, 511, 5, 106, 54, 2, 509, 511, 7, 105, 2,
	2, 510, 498, 3, 2, 2, 2, 510, 499, 3, 2, 2, 2, 510, 500, 3, 2, 2, 2, 510,
	501, 3, 2, 2, 2, 510, 502, 3, 2, 2, 2, 510, 503, 3, 2, 2, 2, 510, 504,
	3, 2, 2, 2, 510, 505, 3, 2, 2, 2, 510, 506, 3, 2, 2, 2, 510, 507, 3, 2,
	2, 2, 510, 508, 3, 2, 2, 2, 510, 509, 3, 2, 2, 2, 511, 99, 3, 2, 2, 2,
	512, 513, 7, 7, 2, 2, 513, 514, 7, 67, 2, 2, 514, 515, 7, 7, 2, 2, 515,
	101, 3, 2, 2, 2, 516, 517, 8, 52, 1, 2, 517, 518, 5, 100, 51, 2, 518, 524,
	3, 2, 2, 2, 519, 520, 12, 3, 2, 2, 520, 521, 7, 68, 2, 2, 521, 523, 5,
	100, 51, 2, 522, 519, 3, 2, 2, 2, 523, 526, 3, 2, 2, 2, 524, 522, 3, 2,
	2, 2, 524, 525, 3, 2, 2, 2, 525, 103, 3, 2, 2, 2, 526, 524, 3, 2, 2, 2,
	527, 531, 7, 47, 2, 2, 528, 529, 7, 64, 2, 2, 529, 530, 7, 7, 2, 2, 530,
	532, 7, 65, 2, 2, 531, 528, 3, 2, 2, 2, 531, 532, 3, 2, 2, 2, 532, 105,
	3, 2, 2, 2, 533, 535, 7, 50, 2, 2, 534, 536, 5, 108, 55, 2, 535, 534, 3,
	2, 2, 2, 536, 537, 3, 2, 2, 2, 537, 535, 3, 2, 2, 2, 537, 538, 3, 2, 2,
	2, 538, 539, 3, 2, 2, 2, 539, 540, 7, 8, 2, 2, 540, 541, 7, 50, 2, 2, 541,
	107, 3, 2, 2, 2, 542, 543, 5, 116, 59, 2, 543, 544, 7, 63, 2, 2, 544, 545,
	5, 98, 50, 2, 545, 109, 3, 2, 2, 2, 546, 547, 7, 12, 2, 2, 547, 548, 5,
	112, 57, 2, 548, 549, 7, 63, 2, 2, 549, 552, 5, 98, 50, 2, 550, 551, 7,
	71, 2, 2, 551, 553, 5, 70, 36, 2, 552, 550, 3, 2, 2, 2, 552, 553, 3, 2,
	2, 2, 553, 560, 3, 2, 2, 2, 554, 555, 7, 12, 2, 2, 555, 556, 5, 112, 57,
	2, 556, 557, 7, 71, 2, 2, 557, 558, 5, 70, 36, 2, 558, 560, 3, 2, 2, 2,
	559, 546, 3, 2, 2, 2, 559, 554, 3, 2, 2, 2, 560, 111, 3, 2, 2, 2, 561,
	562, 7, 105, 2, 2, 562, 113, 3, 2, 2, 2, 563, 564, 7, 48, 2, 2, 564, 565,
	5, 100, 51, 2, 565, 566, 7, 9, 2, 2, 566, 567, 5, 98, 50, 2, 567, 573,
	3, 2, 2, 2, 568, 569, 7, 48, 2, 2, 569, 570, 5, 100, 51, 2, 570, 571, 7,
	68, 2, 2, 571, 573, 3, 2, 2, 2, 572, 563, 3, 2, 2, 2, 572, 568, 3, 2, 2,
	2, 573, 115, 3, 2, 2, 2, 574, 575, 8, 59, 1, 2, 575, 576, 7, 105, 2, 2,
	576, 582, 3, 2, 2, 2, 577, 578, 12, 3, 2, 2, 578, 579, 7, 68, 2, 2, 579,
	581, 7, 105, 2, 2, 580, 577, 3, 2, 2, 2, 581, 584, 3, 2, 2, 2, 582, 580,
	3, 2, 2, 2, 582, 583, 3, 2, 2, 2, 583, 117, 3, 2, 2, 2, 584, 582, 3, 2,
	2, 2, 585, 586, 9, 8, 2, 2, 586, 119, 3, 2, 2, 2, 587, 588, 9, 9, 2, 2,
	588, 121, 3, 2, 2, 2, 57, 125, 129, 135, 138, 144, 147, 150, 156, 160,
	174, 183, 192, 195, 207, 217, 221, 226, 232, 238, 243, 250, 260, 266, 270,
	278, 288, 302, 307, 316, 318, 328, 333, 339, 343, 356, 372, 377, 383, 393,
	404, 415, 428, 435, 447, 466, 479, 486, 510, 524, 531, 537, 552, 559, 572,
	582,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'unqualified'", "'pervasive'", "'opaque'", "", "", "'end'", "'of'",
	"'to'", "'type'", "'var'", "", "", "'class'", "'process'", "'for'", "'loop'",
	"'exit'", "'if'", "'else'", "'elsif'", "'case'", "'assert'", "'begin'",
	"'return'", "'result'", "'new'", "'free'", "'tag'", "'fork'", "'signal'",
	"'wait'", "'pause'", "'quit'", "'unchecked'", "'checked'", "'export'",
	"'import'", "'inherit'", "'implement'", "'by'", "'put'", "'int'", "'real'",
	"'boolean'", "'string'", "'array'", "'set'", "'record'", "'union'", "'pointer'",
	"'nat'", "'intn'", "'natn'", "'realn'", "'char'", "'deferred'", "'forward'",
	"'body'", "'not'", "'^'", "':'", "'('", "')'", "'.'", "'..'", "','", "'#'",
	"'->'", "':='", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'", "'rem'",
	"'**'", "'<'", "'>'", "'='", "'<='", "'>='", "'not='", "'and'", "'or'",
	"'=>'", "'in'", "'shr'", "'shl'", "'xor'", "'+='", "'-='", "'*='", "'/='",
	"'div='", "'mod='", "'and='", "'or='", "'=>='", "'shr='", "'shl='", "'xor='",
}
var symbolicNames = []string{
	"", "", "", "", "STRING_LITERAL", "INTEGER_LITERAL", "END", "OF", "TO",
	"TYPE", "VAR", "PROCEDURE", "FUNCTION", "CLASS", "PROCESS", "FOR", "LOOP",
	"EXIT", "IF", "ELSE", "ELSIF", "CASE", "ASSERT", "BEGIN", "RETURN", "RESULT",
	"NEW", "FREE", "TAG", "FORK", "SIGNAL", "WAIT", "PAUSE", "QUIT", "UNCHECKED",
	"CHECKED", "EXPORT", "IMPORT", "INHERIT", "IMPLEMENT", "BY", "PUT", "INT",
	"REAL", "BOOLEAN", "STRING", "ARRAY", "SET", "RECORD", "UNION", "POINTER",
	"NAT", "INTN", "NATN", "REALN", "CHAR", "DEFERRED", "FORWARD", "BODY",
	"NOT", "CARET", "COLON", "L_PAREN", "R_PAREN", "DOT", "RANGE", "COMMA",
	"CHEAT", "DEREFERENCE", "ASSIGNMENT", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
	"DIV", "MOD", "REM", "EXP", "LESSTHAN", "GREATERTHAN", "EQUAL", "LESSTHANEQUAL",
	"GREATERTHANEQUAL", "NOTEQUAL", "AND", "OR", "IMPLIES", "IN", "SHR", "SHL",
	"XOR", "PLUSEQUALS", "MINUSEQUALS", "MULTIPLYEQUALS", "DIVIDEEQUALS", "DIVEQUALS",
	"MODEQUALS", "ANDEQUALS", "OREQUALS", "IMPLIESEQUALS", "SHREQUALS", "SHLEQUALS",
	"XOREQUALS", "IDENTIFIER", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"program", "topLevel", "procHeader", "funcHeader", "subprogramHeader",
	"subprogramDeclaration", "subprogramPrefix", "classDeclaration", "classBody",
	"exportStatement", "exportListItem", "exportList", "inheritStatement",
	"implementStatement", "idOrFileItem", "howExport", "paramDeclaration",
	"paramDeclarationList", "procBody", "statementOrDeclaration", "primaryExpression",
	"argumentList", "exponentialExpression", "pointerExpression", "postfixExpression",
	"prefixExpression", "multiplicativeExpression", "multiplicativeOperator",
	"additiveExpression", "comparativeExpression", "notExpression", "andExpression",
	"orExpression", "impliesExpression", "expression", "expressionList", "declaration",
	"statements", "assignmentStatement", "putStatement", "putItem", "putItemList",
	"beginStatement", "resultStatement", "newStatement", "freeStatement", "forkStatement",
	"typeDeclaration", "typeSpec", "indexType", "indexTypeList", "stringType",
	"recordType", "recordField", "variableDeclaration", "variableIdentifier",
	"arrayDeclaration", "identifierList", "literal", "comment",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TuringParser struct {
	*antlr.BaseParser
}

func NewTuringParser(input antlr.TokenStream) *TuringParser {
	this := new(TuringParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Turing.g4"

	return this
}

// TuringParser tokens.
const (
	TuringParserEOF              = antlr.TokenEOF
	TuringParserT__0             = 1
	TuringParserT__1             = 2
	TuringParserT__2             = 3
	TuringParserSTRING_LITERAL   = 4
	TuringParserINTEGER_LITERAL  = 5
	TuringParserEND              = 6
	TuringParserOF               = 7
	TuringParserTO               = 8
	TuringParserTYPE             = 9
	TuringParserVAR              = 10
	TuringParserPROCEDURE        = 11
	TuringParserFUNCTION         = 12
	TuringParserCLASS            = 13
	TuringParserPROCESS          = 14
	TuringParserFOR              = 15
	TuringParserLOOP             = 16
	TuringParserEXIT             = 17
	TuringParserIF               = 18
	TuringParserELSE             = 19
	TuringParserELSIF            = 20
	TuringParserCASE             = 21
	TuringParserASSERT           = 22
	TuringParserBEGIN            = 23
	TuringParserRETURN           = 24
	TuringParserRESULT           = 25
	TuringParserNEW              = 26
	TuringParserFREE             = 27
	TuringParserTAG              = 28
	TuringParserFORK             = 29
	TuringParserSIGNAL           = 30
	TuringParserWAIT             = 31
	TuringParserPAUSE            = 32
	TuringParserQUIT             = 33
	TuringParserUNCHECKED        = 34
	TuringParserCHECKED          = 35
	TuringParserEXPORT           = 36
	TuringParserIMPORT           = 37
	TuringParserINHERIT          = 38
	TuringParserIMPLEMENT        = 39
	TuringParserBY               = 40
	TuringParserPUT              = 41
	TuringParserINT              = 42
	TuringParserREAL             = 43
	TuringParserBOOLEAN          = 44
	TuringParserSTRING           = 45
	TuringParserARRAY            = 46
	TuringParserSET              = 47
	TuringParserRECORD           = 48
	TuringParserUNION            = 49
	TuringParserPOINTER          = 50
	TuringParserNAT              = 51
	TuringParserINTN             = 52
	TuringParserNATN             = 53
	TuringParserREALN            = 54
	TuringParserCHAR             = 55
	TuringParserDEFERRED         = 56
	TuringParserFORWARD          = 57
	TuringParserBODY             = 58
	TuringParserNOT              = 59
	TuringParserCARET            = 60
	TuringParserCOLON            = 61
	TuringParserL_PAREN          = 62
	TuringParserR_PAREN          = 63
	TuringParserDOT              = 64
	TuringParserRANGE            = 65
	TuringParserCOMMA            = 66
	TuringParserCHEAT            = 67
	TuringParserDEREFERENCE      = 68
	TuringParserASSIGNMENT       = 69
	TuringParserPLUS             = 70
	TuringParserMINUS            = 71
	TuringParserMULTIPLY         = 72
	TuringParserDIVIDE           = 73
	TuringParserDIV              = 74
	TuringParserMOD              = 75
	TuringParserREM              = 76
	TuringParserEXP              = 77
	TuringParserLESSTHAN         = 78
	TuringParserGREATERTHAN      = 79
	TuringParserEQUAL            = 80
	TuringParserLESSTHANEQUAL    = 81
	TuringParserGREATERTHANEQUAL = 82
	TuringParserNOTEQUAL         = 83
	TuringParserAND              = 84
	TuringParserOR               = 85
	TuringParserIMPLIES          = 86
	TuringParserIN               = 87
	TuringParserSHR              = 88
	TuringParserSHL              = 89
	TuringParserXOR              = 90
	TuringParserPLUSEQUALS       = 91
	TuringParserMINUSEQUALS      = 92
	TuringParserMULTIPLYEQUALS   = 93
	TuringParserDIVIDEEQUALS     = 94
	TuringParserDIVEQUALS        = 95
	TuringParserMODEQUALS        = 96
	TuringParserANDEQUALS        = 97
	TuringParserOREQUALS         = 98
	TuringParserIMPLIESEQUALS    = 99
	TuringParserSHREQUALS        = 100
	TuringParserSHLEQUALS        = 101
	TuringParserXOREQUALS        = 102
	TuringParserIDENTIFIER       = 103
	TuringParserWHITESPACE       = 104
	TuringParserBLOCK_COMMENT    = 105
	TuringParserLINE_COMMENT     = 106
)

// TuringParser rules.
const (
	TuringParserRULE_program                  = 0
	TuringParserRULE_topLevel                 = 1
	TuringParserRULE_procHeader               = 2
	TuringParserRULE_funcHeader               = 3
	TuringParserRULE_subprogramHeader         = 4
	TuringParserRULE_subprogramDeclaration    = 5
	TuringParserRULE_subprogramPrefix         = 6
	TuringParserRULE_classDeclaration         = 7
	TuringParserRULE_classBody                = 8
	TuringParserRULE_exportStatement          = 9
	TuringParserRULE_exportListItem           = 10
	TuringParserRULE_exportList               = 11
	TuringParserRULE_inheritStatement         = 12
	TuringParserRULE_implementStatement       = 13
	TuringParserRULE_idOrFileItem             = 14
	TuringParserRULE_howExport                = 15
	TuringParserRULE_paramDeclaration         = 16
	TuringParserRULE_paramDeclarationList     = 17
	TuringParserRULE_procBody                 = 18
	TuringParserRULE_statementOrDeclaration   = 19
	TuringParserRULE_primaryExpression        = 20
	TuringParserRULE_argumentList             = 21
	TuringParserRULE_exponentialExpression    = 22
	TuringParserRULE_pointerExpression        = 23
	TuringParserRULE_postfixExpression        = 24
	TuringParserRULE_prefixExpression         = 25
	TuringParserRULE_multiplicativeExpression = 26
	TuringParserRULE_multiplicativeOperator   = 27
	TuringParserRULE_additiveExpression       = 28
	TuringParserRULE_comparativeExpression    = 29
	TuringParserRULE_notExpression            = 30
	TuringParserRULE_andExpression            = 31
	TuringParserRULE_orExpression             = 32
	TuringParserRULE_impliesExpression        = 33
	TuringParserRULE_expression               = 34
	TuringParserRULE_expressionList           = 35
	TuringParserRULE_declaration              = 36
	TuringParserRULE_statements               = 37
	TuringParserRULE_assignmentStatement      = 38
	TuringParserRULE_putStatement             = 39
	TuringParserRULE_putItem                  = 40
	TuringParserRULE_putItemList              = 41
	TuringParserRULE_beginStatement           = 42
	TuringParserRULE_resultStatement          = 43
	TuringParserRULE_newStatement             = 44
	TuringParserRULE_freeStatement            = 45
	TuringParserRULE_forkStatement            = 46
	TuringParserRULE_typeDeclaration          = 47
	TuringParserRULE_typeSpec                 = 48
	TuringParserRULE_indexType                = 49
	TuringParserRULE_indexTypeList            = 50
	TuringParserRULE_stringType               = 51
	TuringParserRULE_recordType               = 52
	TuringParserRULE_recordField              = 53
	TuringParserRULE_variableDeclaration      = 54
	TuringParserRULE_variableIdentifier       = 55
	TuringParserRULE_arrayDeclaration         = 56
	TuringParserRULE_identifierList           = 57
	TuringParserRULE_literal                  = 58
	TuringParserRULE_comment                  = 59
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllTopLevel() []ITopLevelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelContext)(nil)).Elem())
	var tst = make([]ITopLevelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelContext)
		}
	}

	return tst
}

func (s *ProgramContext) TopLevel(i int) ITopLevelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *TuringParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TuringParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserCLASS)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(TuringParserPUT-41))|(1<<(TuringParserARRAY-41))|(1<<(TuringParserDEFERRED-41))|(1<<(TuringParserFORWARD-41))|(1<<(TuringParserBODY-41))|(1<<(TuringParserNOT-41))|(1<<(TuringParserCARET-41))|(1<<(TuringParserL_PAREN-41))|(1<<(TuringParserCHEAT-41))|(1<<(TuringParserPLUS-41))|(1<<(TuringParserMINUS-41)))) != 0) || _la == TuringParserIDENTIFIER {
		{
			p.SetState(120)
			p.TopLevel()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopLevelContext is an interface to support dynamic dispatch.
type ITopLevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelContext differentiates from other interfaces.
	IsTopLevelContext()
}

type TopLevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelContext() *TopLevelContext {
	var p = new(TopLevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_topLevel
	return p
}

func (*TopLevelContext) IsTopLevelContext() {}

func NewTopLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelContext {
	var p = new(TopLevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_topLevel

	return p
}

func (s *TopLevelContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelContext) StatementOrDeclaration() IStatementOrDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementOrDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementOrDeclarationContext)
}

func (s *TopLevelContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *TopLevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterTopLevel(s)
	}
}

func (s *TopLevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitTopLevel(s)
	}
}

func (p *TuringParser) TopLevel() (localctx ITopLevelContext) {
	localctx = NewTopLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TuringParserRULE_topLevel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.StatementOrDeclaration()
		}

	case TuringParserCLASS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.ClassDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProcHeaderContext is an interface to support dynamic dispatch.
type IProcHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcHeaderContext differentiates from other interfaces.
	IsProcHeaderContext()
}

type ProcHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcHeaderContext() *ProcHeaderContext {
	var p = new(ProcHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_procHeader
	return p
}

func (*ProcHeaderContext) IsProcHeaderContext() {}

func NewProcHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcHeaderContext {
	var p = new(ProcHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_procHeader

	return p
}

func (s *ProcHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcHeaderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *ProcHeaderContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(TuringParserPROCEDURE, 0)
}

func (s *ProcHeaderContext) PROCESS() antlr.TerminalNode {
	return s.GetToken(TuringParserPROCESS, 0)
}

func (s *ProcHeaderContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *ProcHeaderContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *ProcHeaderContext) ParamDeclarationList() IParamDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamDeclarationListContext)
}

func (s *ProcHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterProcHeader(s)
	}
}

func (s *ProcHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitProcHeader(s)
	}
}

func (p *TuringParser) ProcHeader() (localctx IProcHeaderContext) {
	localctx = NewProcHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TuringParserRULE_procHeader)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TuringParserPROCEDURE || _la == TuringParserPROCESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(130)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(131)
			p.Match(TuringParserL_PAREN)
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS))) != 0) || (((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(TuringParserDEFERRED-56))|(1<<(TuringParserFORWARD-56))|(1<<(TuringParserBODY-56)))) != 0) || _la == TuringParserIDENTIFIER {
			{
				p.SetState(132)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(135)
			p.Match(TuringParserR_PAREN)
		}

	}

	return localctx
}

// IFuncHeaderContext is an interface to support dynamic dispatch.
type IFuncHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncHeaderContext differentiates from other interfaces.
	IsFuncHeaderContext()
}

type FuncHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncHeaderContext() *FuncHeaderContext {
	var p = new(FuncHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_funcHeader
	return p
}

func (*FuncHeaderContext) IsFuncHeaderContext() {}

func NewFuncHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncHeaderContext {
	var p = new(FuncHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_funcHeader

	return p
}

func (s *FuncHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncHeaderContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(TuringParserFUNCTION, 0)
}

func (s *FuncHeaderContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TuringParserIDENTIFIER)
}

func (s *FuncHeaderContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, i)
}

func (s *FuncHeaderContext) COLON() antlr.TerminalNode {
	return s.GetToken(TuringParserCOLON, 0)
}

func (s *FuncHeaderContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FuncHeaderContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *FuncHeaderContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *FuncHeaderContext) ParamDeclarationList() IParamDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamDeclarationListContext)
}

func (s *FuncHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterFuncHeader(s)
	}
}

func (s *FuncHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitFuncHeader(s)
	}
}

func (p *TuringParser) FuncHeader() (localctx IFuncHeaderContext) {
	localctx = NewFuncHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TuringParserRULE_funcHeader)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(TuringParserFUNCTION)
	}
	{
		p.SetState(139)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TuringParserL_PAREN {
		{
			p.SetState(140)
			p.Match(TuringParserL_PAREN)
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS))) != 0) || (((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(TuringParserDEFERRED-56))|(1<<(TuringParserFORWARD-56))|(1<<(TuringParserBODY-56)))) != 0) || _la == TuringParserIDENTIFIER {
			{
				p.SetState(141)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(144)
			p.Match(TuringParserR_PAREN)
		}

	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TuringParserIDENTIFIER {
		{
			p.SetState(147)
			p.Match(TuringParserIDENTIFIER)
		}

	}
	{
		p.SetState(150)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(151)
		p.TypeSpec()
	}

	return localctx
}

// ISubprogramHeaderContext is an interface to support dynamic dispatch.
type ISubprogramHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubprogramHeaderContext differentiates from other interfaces.
	IsSubprogramHeaderContext()
}

type SubprogramHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubprogramHeaderContext() *SubprogramHeaderContext {
	var p = new(SubprogramHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_subprogramHeader
	return p
}

func (*SubprogramHeaderContext) IsSubprogramHeaderContext() {}

func NewSubprogramHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubprogramHeaderContext {
	var p = new(SubprogramHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_subprogramHeader

	return p
}

func (s *SubprogramHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *SubprogramHeaderContext) FuncHeader() IFuncHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncHeaderContext)
}

func (s *SubprogramHeaderContext) ProcHeader() IProcHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcHeaderContext)
}

func (s *SubprogramHeaderContext) SubprogramPrefix() ISubprogramPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubprogramPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubprogramPrefixContext)
}

func (s *SubprogramHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubprogramHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubprogramHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterSubprogramHeader(s)
	}
}

func (s *SubprogramHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitSubprogramHeader(s)
	}
}

func (p *TuringParser) SubprogramHeader() (localctx ISubprogramHeaderContext) {
	localctx = NewSubprogramHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TuringParserRULE_subprogramHeader)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(TuringParserDEFERRED-56))|(1<<(TuringParserFORWARD-56))|(1<<(TuringParserBODY-56)))) != 0 {
		{
			p.SetState(153)
			p.SubprogramPrefix()
		}

	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserFUNCTION:
		{
			p.SetState(156)
			p.FuncHeader()
		}

	case TuringParserPROCEDURE, TuringParserPROCESS:
		{
			p.SetState(157)
			p.ProcHeader()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubprogramDeclarationContext is an interface to support dynamic dispatch.
type ISubprogramDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubprogramDeclarationContext differentiates from other interfaces.
	IsSubprogramDeclarationContext()
}

type SubprogramDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubprogramDeclarationContext() *SubprogramDeclarationContext {
	var p = new(SubprogramDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_subprogramDeclaration
	return p
}

func (*SubprogramDeclarationContext) IsSubprogramDeclarationContext() {}

func NewSubprogramDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubprogramDeclarationContext {
	var p = new(SubprogramDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_subprogramDeclaration

	return p
}

func (s *SubprogramDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SubprogramDeclarationContext) SubprogramHeader() ISubprogramHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubprogramHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubprogramHeaderContext)
}

func (s *SubprogramDeclarationContext) ProcBody() IProcBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcBodyContext)
}

func (s *SubprogramDeclarationContext) END() antlr.TerminalNode {
	return s.GetToken(TuringParserEND, 0)
}

func (s *SubprogramDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *SubprogramDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubprogramDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubprogramDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterSubprogramDeclaration(s)
	}
}

func (s *SubprogramDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitSubprogramDeclaration(s)
	}
}

func (p *TuringParser) SubprogramDeclaration() (localctx ISubprogramDeclarationContext) {
	localctx = NewSubprogramDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TuringParserRULE_subprogramDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.SubprogramHeader()
	}
	{
		p.SetState(161)
		p.ProcBody()
	}
	{
		p.SetState(162)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(163)
		p.Match(TuringParserIDENTIFIER)
	}

	return localctx
}

// ISubprogramPrefixContext is an interface to support dynamic dispatch.
type ISubprogramPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubprogramPrefixContext differentiates from other interfaces.
	IsSubprogramPrefixContext()
}

type SubprogramPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubprogramPrefixContext() *SubprogramPrefixContext {
	var p = new(SubprogramPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_subprogramPrefix
	return p
}

func (*SubprogramPrefixContext) IsSubprogramPrefixContext() {}

func NewSubprogramPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubprogramPrefixContext {
	var p = new(SubprogramPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_subprogramPrefix

	return p
}

func (s *SubprogramPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SubprogramPrefixContext) DEFERRED() antlr.TerminalNode {
	return s.GetToken(TuringParserDEFERRED, 0)
}

func (s *SubprogramPrefixContext) BODY() antlr.TerminalNode {
	return s.GetToken(TuringParserBODY, 0)
}

func (s *SubprogramPrefixContext) FORWARD() antlr.TerminalNode {
	return s.GetToken(TuringParserFORWARD, 0)
}

func (s *SubprogramPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubprogramPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubprogramPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterSubprogramPrefix(s)
	}
}

func (s *SubprogramPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitSubprogramPrefix(s)
	}
}

func (p *TuringParser) SubprogramPrefix() (localctx ISubprogramPrefixContext) {
	localctx = NewSubprogramPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TuringParserRULE_subprogramPrefix)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(TuringParserDEFERRED-56))|(1<<(TuringParserFORWARD-56))|(1<<(TuringParserBODY-56)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(TuringParserCLASS, 0)
}

func (s *ClassDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TuringParserIDENTIFIER)
}

func (s *ClassDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, i)
}

func (s *ClassDeclarationContext) END() antlr.TerminalNode {
	return s.GetToken(TuringParserEND, 0)
}

func (s *ClassDeclarationContext) AllClassBody() []IClassBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassBodyContext)(nil)).Elem())
	var tst = make([]IClassBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassBodyContext)
		}
	}

	return tst
}

func (s *ClassDeclarationContext) ClassBody(i int) IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (p *TuringParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TuringParserRULE_classDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(TuringParserCLASS)
	}
	{
		p.SetState(168)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(TuringParserEXPORT-36))|(1<<(TuringParserINHERIT-36))|(1<<(TuringParserIMPLEMENT-36))|(1<<(TuringParserPUT-36))|(1<<(TuringParserARRAY-36))|(1<<(TuringParserDEFERRED-36))|(1<<(TuringParserFORWARD-36))|(1<<(TuringParserBODY-36))|(1<<(TuringParserNOT-36))|(1<<(TuringParserCARET-36))|(1<<(TuringParserL_PAREN-36))|(1<<(TuringParserCHEAT-36)))) != 0) || _la == TuringParserPLUS || _la == TuringParserMINUS || _la == TuringParserIDENTIFIER {
		{
			p.SetState(169)
			p.ClassBody()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(175)
		p.Match(TuringParserIDENTIFIER)
	}

	return localctx
}

// IClassBodyContext is an interface to support dynamic dispatch.
type IClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyContext differentiates from other interfaces.
	IsClassBodyContext()
}

type ClassBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyContext() *ClassBodyContext {
	var p = new(ClassBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_classBody
	return p
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_classBody

	return p
}

func (s *ClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyContext) ExportStatement() IExportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportStatementContext)
}

func (s *ClassBodyContext) InheritStatement() IInheritStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInheritStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInheritStatementContext)
}

func (s *ClassBodyContext) ImplementStatement() IImplementStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementStatementContext)
}

func (s *ClassBodyContext) StatementOrDeclaration() IStatementOrDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementOrDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementOrDeclarationContext)
}

func (s *ClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (p *TuringParser) ClassBody() (localctx IClassBodyContext) {
	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TuringParserRULE_classBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserEXPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.ExportStatement()
		}

	case TuringParserINHERIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.InheritStatement()
		}

	case TuringParserIMPLEMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(179)
			p.ImplementStatement()
		}

	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(180)
			p.StatementOrDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExportStatementContext is an interface to support dynamic dispatch.
type IExportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportStatementContext differentiates from other interfaces.
	IsExportStatementContext()
}

type ExportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportStatementContext() *ExportStatementContext {
	var p = new(ExportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_exportStatement
	return p
}

func (*ExportStatementContext) IsExportStatementContext() {}

func NewExportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportStatementContext {
	var p = new(ExportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_exportStatement

	return p
}

func (s *ExportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportStatementContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(TuringParserEXPORT, 0)
}

func (s *ExportStatementContext) ExportList() IExportListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportListContext)
}

func (s *ExportStatementContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *ExportStatementContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *ExportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterExportStatement(s)
	}
}

func (s *ExportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitExportStatement(s)
	}
}

func (p *TuringParser) ExportStatement() (localctx IExportStatementContext) {
	localctx = NewExportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TuringParserRULE_exportStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(TuringParserEXPORT)
		}
		{
			p.SetState(184)
			p.exportList(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(TuringParserEXPORT)
		}
		{
			p.SetState(186)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(187)
			p.exportList(0)
		}
		{
			p.SetState(188)
			p.Match(TuringParserR_PAREN)
		}

	}

	return localctx
}

// IExportListItemContext is an interface to support dynamic dispatch.
type IExportListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportListItemContext differentiates from other interfaces.
	IsExportListItemContext()
}

type ExportListItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportListItemContext() *ExportListItemContext {
	var p = new(ExportListItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_exportListItem
	return p
}

func (*ExportListItemContext) IsExportListItemContext() {}

func NewExportListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportListItemContext {
	var p = new(ExportListItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_exportListItem

	return p
}

func (s *ExportListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportListItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *ExportListItemContext) HowExport() IHowExportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHowExportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHowExportContext)
}

func (s *ExportListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterExportListItem(s)
	}
}

func (s *ExportListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitExportListItem(s)
	}
}

func (p *TuringParser) ExportListItem() (localctx IExportListItemContext) {
	localctx = NewExportListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TuringParserRULE_exportListItem)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserT__0)|(1<<TuringParserT__1)|(1<<TuringParserT__2)|(1<<TuringParserVAR))) != 0 {
		{
			p.SetState(192)
			p.HowExport()
		}

	}
	{
		p.SetState(195)
		p.Match(TuringParserIDENTIFIER)
	}

	return localctx
}

// IExportListContext is an interface to support dynamic dispatch.
type IExportListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportListContext differentiates from other interfaces.
	IsExportListContext()
}

type ExportListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportListContext() *ExportListContext {
	var p = new(ExportListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_exportList
	return p
}

func (*ExportListContext) IsExportListContext() {}

func NewExportListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportListContext {
	var p = new(ExportListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_exportList

	return p
}

func (s *ExportListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportListContext) ExportListItem() IExportListItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportListItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportListItemContext)
}

func (s *ExportListContext) ExportList() IExportListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportListContext)
}

func (s *ExportListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *ExportListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterExportList(s)
	}
}

func (s *ExportListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitExportList(s)
	}
}

func (p *TuringParser) ExportList() (localctx IExportListContext) {
	return p.exportList(0)
}

func (p *TuringParser) exportList(_p int) (localctx IExportListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExportListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExportListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, TuringParserRULE_exportList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.ExportListItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExportListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_exportList)
			p.SetState(200)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(201)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(202)
				p.ExportListItem()
			}

		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IInheritStatementContext is an interface to support dynamic dispatch.
type IInheritStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInheritStatementContext differentiates from other interfaces.
	IsInheritStatementContext()
}

type InheritStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInheritStatementContext() *InheritStatementContext {
	var p = new(InheritStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_inheritStatement
	return p
}

func (*InheritStatementContext) IsInheritStatementContext() {}

func NewInheritStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InheritStatementContext {
	var p = new(InheritStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_inheritStatement

	return p
}

func (s *InheritStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InheritStatementContext) INHERIT() antlr.TerminalNode {
	return s.GetToken(TuringParserINHERIT, 0)
}

func (s *InheritStatementContext) IdOrFileItem() IIdOrFileItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdOrFileItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdOrFileItemContext)
}

func (s *InheritStatementContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *InheritStatementContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *InheritStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InheritStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterInheritStatement(s)
	}
}

func (s *InheritStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitInheritStatement(s)
	}
}

func (p *TuringParser) InheritStatement() (localctx IInheritStatementContext) {
	localctx = NewInheritStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TuringParserRULE_inheritStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.Match(TuringParserINHERIT)
		}
		{
			p.SetState(209)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.Match(TuringParserINHERIT)
		}
		{
			p.SetState(211)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(212)
			p.IdOrFileItem()
		}
		{
			p.SetState(213)
			p.Match(TuringParserR_PAREN)
		}

	}

	return localctx
}

// IImplementStatementContext is an interface to support dynamic dispatch.
type IImplementStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplementStatementContext differentiates from other interfaces.
	IsImplementStatementContext()
}

type ImplementStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementStatementContext() *ImplementStatementContext {
	var p = new(ImplementStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_implementStatement
	return p
}

func (*ImplementStatementContext) IsImplementStatementContext() {}

func NewImplementStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementStatementContext {
	var p = new(ImplementStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_implementStatement

	return p
}

func (s *ImplementStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementStatementContext) IMPLEMENT() antlr.TerminalNode {
	return s.GetToken(TuringParserIMPLEMENT, 0)
}

func (s *ImplementStatementContext) IdOrFileItem() IIdOrFileItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdOrFileItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdOrFileItemContext)
}

func (s *ImplementStatementContext) BY() antlr.TerminalNode {
	return s.GetToken(TuringParserBY, 0)
}

func (s *ImplementStatementContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *ImplementStatementContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *ImplementStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterImplementStatement(s)
	}
}

func (s *ImplementStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitImplementStatement(s)
	}
}

func (p *TuringParser) ImplementStatement() (localctx IImplementStatementContext) {
	localctx = NewImplementStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TuringParserRULE_implementStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(TuringParserIMPLEMENT)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserBY {
			{
				p.SetState(218)
				p.Match(TuringParserBY)
			}

		}
		{
			p.SetState(221)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(TuringParserIMPLEMENT)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserBY {
			{
				p.SetState(223)
				p.Match(TuringParserBY)
			}

		}
		{
			p.SetState(226)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(227)
			p.IdOrFileItem()
		}
		{
			p.SetState(228)
			p.Match(TuringParserR_PAREN)
		}

	}

	return localctx
}

// IIdOrFileItemContext is an interface to support dynamic dispatch.
type IIdOrFileItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdOrFileItemContext differentiates from other interfaces.
	IsIdOrFileItemContext()
}

type IdOrFileItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdOrFileItemContext() *IdOrFileItemContext {
	var p = new(IdOrFileItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_idOrFileItem
	return p
}

func (*IdOrFileItemContext) IsIdOrFileItemContext() {}

func NewIdOrFileItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdOrFileItemContext {
	var p = new(IdOrFileItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_idOrFileItem

	return p
}

func (s *IdOrFileItemContext) GetParser() antlr.Parser { return s.parser }

func (s *IdOrFileItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *IdOrFileItemContext) IN() antlr.TerminalNode {
	return s.GetToken(TuringParserIN, 0)
}

func (s *IdOrFileItemContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(TuringParserSTRING_LITERAL, 0)
}

func (s *IdOrFileItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdOrFileItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdOrFileItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterIdOrFileItem(s)
	}
}

func (s *IdOrFileItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitIdOrFileItem(s)
	}
}

func (p *TuringParser) IdOrFileItem() (localctx IIdOrFileItemContext) {
	localctx = NewIdOrFileItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TuringParserRULE_idOrFileItem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(TuringParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(234)
			p.Match(TuringParserIN)
		}
		{
			p.SetState(235)
			p.Match(TuringParserSTRING_LITERAL)
		}

	}

	return localctx
}

// IHowExportContext is an interface to support dynamic dispatch.
type IHowExportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHowExportContext differentiates from other interfaces.
	IsHowExportContext()
}

type HowExportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHowExportContext() *HowExportContext {
	var p = new(HowExportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_howExport
	return p
}

func (*HowExportContext) IsHowExportContext() {}

func NewHowExportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HowExportContext {
	var p = new(HowExportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_howExport

	return p
}

func (s *HowExportContext) GetParser() antlr.Parser { return s.parser }

func (s *HowExportContext) VAR() antlr.TerminalNode {
	return s.GetToken(TuringParserVAR, 0)
}

func (s *HowExportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HowExportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HowExportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterHowExport(s)
	}
}

func (s *HowExportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitHowExport(s)
	}
}

func (p *TuringParser) HowExport() (localctx IHowExportContext) {
	localctx = NewHowExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TuringParserRULE_howExport)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserT__0)|(1<<TuringParserT__1)|(1<<TuringParserT__2)|(1<<TuringParserVAR))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParamDeclarationContext is an interface to support dynamic dispatch.
type IParamDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamDeclarationContext differentiates from other interfaces.
	IsParamDeclarationContext()
}

type ParamDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamDeclarationContext() *ParamDeclarationContext {
	var p = new(ParamDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_paramDeclaration
	return p
}

func (*ParamDeclarationContext) IsParamDeclarationContext() {}

func NewParamDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclarationContext {
	var p = new(ParamDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_paramDeclaration

	return p
}

func (s *ParamDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclarationContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ParamDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(TuringParserCOLON, 0)
}

func (s *ParamDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ParamDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(TuringParserVAR, 0)
}

func (s *ParamDeclarationContext) SubprogramHeader() ISubprogramHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubprogramHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubprogramHeaderContext)
}

func (s *ParamDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterParamDeclaration(s)
	}
}

func (s *ParamDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitParamDeclaration(s)
	}
}

func (p *TuringParser) ParamDeclaration() (localctx IParamDeclarationContext) {
	localctx = NewParamDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TuringParserRULE_paramDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserVAR, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserVAR {
			{
				p.SetState(240)
				p.Match(TuringParserVAR)
			}

		}
		{
			p.SetState(243)
			p.identifierList(0)
		}
		{
			p.SetState(244)
			p.Match(TuringParserCOLON)
		}
		{
			p.SetState(245)
			p.TypeSpec()
		}

	case TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.SubprogramHeader()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParamDeclarationListContext is an interface to support dynamic dispatch.
type IParamDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamDeclarationListContext differentiates from other interfaces.
	IsParamDeclarationListContext()
}

type ParamDeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamDeclarationListContext() *ParamDeclarationListContext {
	var p = new(ParamDeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_paramDeclarationList
	return p
}

func (*ParamDeclarationListContext) IsParamDeclarationListContext() {}

func NewParamDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclarationListContext {
	var p = new(ParamDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_paramDeclarationList

	return p
}

func (s *ParamDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclarationListContext) ParamDeclaration() IParamDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamDeclarationContext)
}

func (s *ParamDeclarationListContext) ParamDeclarationList() IParamDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamDeclarationListContext)
}

func (s *ParamDeclarationListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *ParamDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterParamDeclarationList(s)
	}
}

func (s *ParamDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitParamDeclarationList(s)
	}
}

func (p *TuringParser) ParamDeclarationList() (localctx IParamDeclarationListContext) {
	return p.paramDeclarationList(0)
}

func (p *TuringParser) paramDeclarationList(_p int) (localctx IParamDeclarationListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParamDeclarationListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParamDeclarationListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, TuringParserRULE_paramDeclarationList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.ParamDeclaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParamDeclarationListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_paramDeclarationList)
			p.SetState(253)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(254)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(255)
				p.ParamDeclaration()
			}

		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IProcBodyContext is an interface to support dynamic dispatch.
type IProcBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcBodyContext differentiates from other interfaces.
	IsProcBodyContext()
}

type ProcBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcBodyContext() *ProcBodyContext {
	var p = new(ProcBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_procBody
	return p
}

func (*ProcBodyContext) IsProcBodyContext() {}

func NewProcBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcBodyContext {
	var p = new(ProcBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_procBody

	return p
}

func (s *ProcBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcBodyContext) AllStatementOrDeclaration() []IStatementOrDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementOrDeclarationContext)(nil)).Elem())
	var tst = make([]IStatementOrDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementOrDeclarationContext)
		}
	}

	return tst
}

func (s *ProcBodyContext) StatementOrDeclaration(i int) IStatementOrDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementOrDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementOrDeclarationContext)
}

func (s *ProcBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterProcBody(s)
	}
}

func (s *ProcBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitProcBody(s)
	}
}

func (p *TuringParser) ProcBody() (localctx IProcBodyContext) {
	localctx = NewProcBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TuringParserRULE_procBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(TuringParserPUT-41))|(1<<(TuringParserARRAY-41))|(1<<(TuringParserDEFERRED-41))|(1<<(TuringParserFORWARD-41))|(1<<(TuringParserBODY-41))|(1<<(TuringParserNOT-41))|(1<<(TuringParserCARET-41))|(1<<(TuringParserL_PAREN-41))|(1<<(TuringParserCHEAT-41))|(1<<(TuringParserPLUS-41))|(1<<(TuringParserMINUS-41)))) != 0) || _la == TuringParserIDENTIFIER {
		{
			p.SetState(261)
			p.StatementOrDeclaration()
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementOrDeclarationContext is an interface to support dynamic dispatch.
type IStatementOrDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementOrDeclarationContext differentiates from other interfaces.
	IsStatementOrDeclarationContext()
}

type StatementOrDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementOrDeclarationContext() *StatementOrDeclarationContext {
	var p = new(StatementOrDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_statementOrDeclaration
	return p
}

func (*StatementOrDeclarationContext) IsStatementOrDeclarationContext() {}

func NewStatementOrDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementOrDeclarationContext {
	var p = new(StatementOrDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_statementOrDeclaration

	return p
}

func (s *StatementOrDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementOrDeclarationContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *StatementOrDeclarationContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementOrDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementOrDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementOrDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterStatementOrDeclaration(s)
	}
}

func (s *StatementOrDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitStatementOrDeclaration(s)
	}
}

func (p *TuringParser) StatementOrDeclaration() (localctx IStatementOrDeclarationContext) {
	localctx = NewStatementOrDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TuringParserRULE_statementOrDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Statements()
		}

	case TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Declaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *PrimaryExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (p *TuringParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TuringParserRULE_primaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(276)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.Literal()
		}

	case TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
			p.Match(TuringParserIDENTIFIER)
		}

	case TuringParserL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(272)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(273)
			p.Expression()
		}
		{
			p.SetState(274)
			p.Match(TuringParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllArgumentList() []IArgumentListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentListContext)(nil)).Elem())
	var tst = make([]IArgumentListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentListContext)
		}
	}

	return tst
}

func (s *ArgumentListContext) ArgumentList(i int) IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *ArgumentListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (p *TuringParser) ArgumentList() (localctx IArgumentListContext) {
	return p.argumentList(0)
}

func (p *TuringParser) argumentList(_p int) (localctx IArgumentListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArgumentListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, TuringParserRULE_argumentList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArgumentListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_argumentList)
			p.SetState(281)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(282)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(283)
				p.argumentList(2)
			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IExponentialExpressionContext is an interface to support dynamic dispatch.
type IExponentialExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExponentialExpressionContext differentiates from other interfaces.
	IsExponentialExpressionContext()
}

type ExponentialExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExponentialExpressionContext() *ExponentialExpressionContext {
	var p = new(ExponentialExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_exponentialExpression
	return p
}

func (*ExponentialExpressionContext) IsExponentialExpressionContext() {}

func NewExponentialExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExponentialExpressionContext {
	var p = new(ExponentialExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_exponentialExpression

	return p
}

func (s *ExponentialExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExponentialExpressionContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem())
	var tst = make([]IPrimaryExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExpressionContext)
		}
	}

	return tst
}

func (s *ExponentialExpressionContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExponentialExpressionContext) EXP() antlr.TerminalNode {
	return s.GetToken(TuringParserEXP, 0)
}

func (s *ExponentialExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentialExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExponentialExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterExponentialExpression(s)
	}
}

func (s *ExponentialExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitExponentialExpression(s)
	}
}

func (p *TuringParser) ExponentialExpression() (localctx IExponentialExpressionContext) {
	localctx = NewExponentialExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TuringParserRULE_exponentialExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.PrimaryExpression()
	}
	{
		p.SetState(290)
		p.Match(TuringParserEXP)
	}
	{
		p.SetState(291)
		p.PrimaryExpression()
	}

	return localctx
}

// IPointerExpressionContext is an interface to support dynamic dispatch.
type IPointerExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointerExpressionContext differentiates from other interfaces.
	IsPointerExpressionContext()
}

type PointerExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerExpressionContext() *PointerExpressionContext {
	var p = new(PointerExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_pointerExpression
	return p
}

func (*PointerExpressionContext) IsPointerExpressionContext() {}

func NewPointerExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerExpressionContext {
	var p = new(PointerExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_pointerExpression

	return p
}

func (s *PointerExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(TuringParserCARET, 0)
}

func (s *PointerExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PointerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterPointerExpression(s)
	}
}

func (s *PointerExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitPointerExpression(s)
	}
}

func (p *TuringParser) PointerExpression() (localctx IPointerExpressionContext) {
	localctx = NewPointerExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TuringParserRULE_pointerExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(TuringParserCARET)
	}
	{
		p.SetState(294)
		p.PrimaryExpression()
	}

	return localctx
}

// IPostfixExpressionContext is an interface to support dynamic dispatch.
type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostfixExpressionContext differentiates from other interfaces.
	IsPostfixExpressionContext()
}

type PostfixExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixExpressionContext() *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_postfixExpression
	return p
}

func (*PostfixExpressionContext) IsPostfixExpressionContext() {}

func NewPostfixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_postfixExpression

	return p
}

func (s *PostfixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PostfixExpressionContext) PointerExpression() IPointerExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerExpressionContext)
}

func (s *PostfixExpressionContext) ExponentialExpression() IExponentialExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExponentialExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExponentialExpressionContext)
}

func (s *PostfixExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *PostfixExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *PostfixExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *PostfixExpressionContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *PostfixExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(TuringParserDOT, 0)
}

func (s *PostfixExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *PostfixExpressionContext) DEREFERENCE() antlr.TerminalNode {
	return s.GetToken(TuringParserDEREFERENCE, 0)
}

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitPostfixExpression(s)
	}
}

func (p *TuringParser) PostfixExpression() (localctx IPostfixExpressionContext) {
	return p.postfixExpression(0)
}

func (p *TuringParser) postfixExpression(_p int) (localctx IPostfixExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPostfixExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPostfixExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, TuringParserRULE_postfixExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(297)
			p.PrimaryExpression()
		}

	case 2:
		{
			p.SetState(298)
			p.PointerExpression()
		}

	case 3:
		{
			p.SetState(299)
			p.ExponentialExpression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(314)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_postfixExpression)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(303)
					p.Match(TuringParserL_PAREN)
				}
				p.SetState(305)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == TuringParserSTRING_LITERAL || _la == TuringParserINTEGER_LITERAL || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(TuringParserNOT-59))|(1<<(TuringParserCARET-59))|(1<<(TuringParserL_PAREN-59))|(1<<(TuringParserCHEAT-59))|(1<<(TuringParserPLUS-59))|(1<<(TuringParserMINUS-59)))) != 0) || _la == TuringParserIDENTIFIER {
					{
						p.SetState(304)
						p.argumentList(0)
					}

				}
				{
					p.SetState(307)
					p.Match(TuringParserR_PAREN)
				}

			case 2:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_postfixExpression)
				p.SetState(308)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(309)
					p.Match(TuringParserDOT)
				}
				{
					p.SetState(310)
					p.Match(TuringParserIDENTIFIER)
				}

			case 3:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_postfixExpression)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(312)
					p.Match(TuringParserDEREFERENCE)
				}
				{
					p.SetState(313)
					p.Match(TuringParserIDENTIFIER)
				}

			}

		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IPrefixExpressionContext is an interface to support dynamic dispatch.
type IPrefixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixExpressionContext differentiates from other interfaces.
	IsPrefixExpressionContext()
}

type PrefixExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixExpressionContext() *PrefixExpressionContext {
	var p = new(PrefixExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_prefixExpression
	return p
}

func (*PrefixExpressionContext) IsPrefixExpressionContext() {}

func NewPrefixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixExpressionContext {
	var p = new(PrefixExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_prefixExpression

	return p
}

func (s *PrefixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *PrefixExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TuringParserPLUS, 0)
}

func (s *PrefixExpressionContext) PrefixExpression() IPrefixExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixExpressionContext)
}

func (s *PrefixExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TuringParserMINUS, 0)
}

func (s *PrefixExpressionContext) CHEAT() antlr.TerminalNode {
	return s.GetToken(TuringParserCHEAT, 0)
}

func (s *PrefixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterPrefixExpression(s)
	}
}

func (s *PrefixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitPrefixExpression(s)
	}
}

func (p *TuringParser) PrefixExpression() (localctx IPrefixExpressionContext) {
	localctx = NewPrefixExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TuringParserRULE_prefixExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(326)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserCARET, TuringParserL_PAREN, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.postfixExpression(0)
		}

	case TuringParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(TuringParserPLUS)
		}
		{
			p.SetState(321)
			p.PrefixExpression()
		}

	case TuringParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(322)
			p.Match(TuringParserMINUS)
		}
		{
			p.SetState(323)
			p.PrefixExpression()
		}

	case TuringParserCHEAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(324)
			p.Match(TuringParserCHEAT)
		}
		{
			p.SetState(325)
			p.PrefixExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *MultiplicativeExpressionContext) PrefixExpression() IPrefixExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeOperator() IMultiplicativeOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeOperatorContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (p *TuringParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	return p.multiplicativeExpression(0)
}

func (p *TuringParser) multiplicativeExpression(_p int) (localctx IMultiplicativeExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, TuringParserRULE_multiplicativeExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(329)
			p.postfixExpression(0)
		}

	case 2:
		{
			p.SetState(330)
			p.PrefixExpression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_multiplicativeExpression)
			p.SetState(333)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(334)
				p.MultiplicativeOperator()
			}
			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(335)
					p.postfixExpression(0)
				}

			case 2:
				{
					p.SetState(336)
					p.PrefixExpression()
				}

			}

		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplicativeOperatorContext is an interface to support dynamic dispatch.
type IMultiplicativeOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeOperatorContext differentiates from other interfaces.
	IsMultiplicativeOperatorContext()
}

type MultiplicativeOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOperatorContext() *MultiplicativeOperatorContext {
	var p = new(MultiplicativeOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_multiplicativeOperator
	return p
}

func (*MultiplicativeOperatorContext) IsMultiplicativeOperatorContext() {}

func NewMultiplicativeOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOperatorContext {
	var p = new(MultiplicativeOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_multiplicativeOperator

	return p
}

func (s *MultiplicativeOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOperatorContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(TuringParserMULTIPLY, 0)
}

func (s *MultiplicativeOperatorContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVIDE, 0)
}

func (s *MultiplicativeOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(TuringParserDIV, 0)
}

func (s *MultiplicativeOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(TuringParserMOD, 0)
}

func (s *MultiplicativeOperatorContext) REM() antlr.TerminalNode {
	return s.GetToken(TuringParserREM, 0)
}

func (s *MultiplicativeOperatorContext) SHR() antlr.TerminalNode {
	return s.GetToken(TuringParserSHR, 0)
}

func (s *MultiplicativeOperatorContext) SHL() antlr.TerminalNode {
	return s.GetToken(TuringParserSHL, 0)
}

func (s *MultiplicativeOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterMultiplicativeOperator(s)
	}
}

func (s *MultiplicativeOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitMultiplicativeOperator(s)
	}
}

func (p *TuringParser) MultiplicativeOperator() (localctx IMultiplicativeOperatorContext) {
	localctx = NewMultiplicativeOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TuringParserRULE_multiplicativeOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(TuringParserMULTIPLY-72))|(1<<(TuringParserDIVIDE-72))|(1<<(TuringParserDIV-72))|(1<<(TuringParserMOD-72))|(1<<(TuringParserREM-72))|(1<<(TuringParserSHR-72))|(1<<(TuringParserSHL-72)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *AdditiveExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TuringParserPLUS, 0)
}

func (s *AdditiveExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TuringParserMINUS, 0)
}

func (s *AdditiveExpressionContext) XOR() antlr.TerminalNode {
	return s.GetToken(TuringParserXOR, 0)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (p *TuringParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	return p.additiveExpression(0)
}

func (p *TuringParser) additiveExpression(_p int) (localctx IAdditiveExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, TuringParserRULE_additiveExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.multiplicativeExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_additiveExpression)
			p.SetState(349)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(350)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(TuringParserPLUS-70))|(1<<(TuringParserMINUS-70))|(1<<(TuringParserXOR-70)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(351)
				p.multiplicativeExpression(0)
			}

		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IComparativeExpressionContext is an interface to support dynamic dispatch.
type IComparativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparativeExpressionContext differentiates from other interfaces.
	IsComparativeExpressionContext()
}

type ComparativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparativeExpressionContext() *ComparativeExpressionContext {
	var p = new(ComparativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_comparativeExpression
	return p
}

func (*ComparativeExpressionContext) IsComparativeExpressionContext() {}

func NewComparativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparativeExpressionContext {
	var p = new(ComparativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_comparativeExpression

	return p
}

func (s *ComparativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparativeExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *ComparativeExpressionContext) ComparativeExpression() IComparativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparativeExpressionContext)
}

func (s *ComparativeExpressionContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(TuringParserLESSTHAN, 0)
}

func (s *ComparativeExpressionContext) GREATERTHAN() antlr.TerminalNode {
	return s.GetToken(TuringParserGREATERTHAN, 0)
}

func (s *ComparativeExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserEQUAL, 0)
}

func (s *ComparativeExpressionContext) LESSTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserLESSTHANEQUAL, 0)
}

func (s *ComparativeExpressionContext) GREATERTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserGREATERTHANEQUAL, 0)
}

func (s *ComparativeExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserNOTEQUAL, 0)
}

func (s *ComparativeExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(TuringParserIN, 0)
}

func (s *ComparativeExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TuringParserNOT, 0)
}

func (s *ComparativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterComparativeExpression(s)
	}
}

func (s *ComparativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitComparativeExpression(s)
	}
}

func (p *TuringParser) ComparativeExpression() (localctx IComparativeExpressionContext) {
	return p.comparativeExpression(0)
}

func (p *TuringParser) comparativeExpression(_p int) (localctx IComparativeExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewComparativeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparativeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, TuringParserRULE_comparativeExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.additiveExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComparativeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_comparativeExpression)
			p.SetState(360)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(370)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case TuringParserLESSTHAN:
				{
					p.SetState(361)
					p.Match(TuringParserLESSTHAN)
				}

			case TuringParserGREATERTHAN:
				{
					p.SetState(362)
					p.Match(TuringParserGREATERTHAN)
				}

			case TuringParserEQUAL:
				{
					p.SetState(363)
					p.Match(TuringParserEQUAL)
				}

			case TuringParserLESSTHANEQUAL:
				{
					p.SetState(364)
					p.Match(TuringParserLESSTHANEQUAL)
				}

			case TuringParserGREATERTHANEQUAL:
				{
					p.SetState(365)
					p.Match(TuringParserGREATERTHANEQUAL)
				}

			case TuringParserNOTEQUAL:
				{
					p.SetState(366)
					p.Match(TuringParserNOTEQUAL)
				}

			case TuringParserIN:
				{
					p.SetState(367)
					p.Match(TuringParserIN)
				}

			case TuringParserNOT:
				{
					p.SetState(368)
					p.Match(TuringParserNOT)
				}
				{
					p.SetState(369)
					p.Match(TuringParserIN)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(372)
				p.additiveExpression(0)
			}

		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}

	return localctx
}

// INotExpressionContext is an interface to support dynamic dispatch.
type INotExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotExpressionContext differentiates from other interfaces.
	IsNotExpressionContext()
}

type NotExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotExpressionContext() *NotExpressionContext {
	var p = new(NotExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_notExpression
	return p
}

func (*NotExpressionContext) IsNotExpressionContext() {}

func NewNotExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_notExpression

	return p
}

func (s *NotExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NotExpressionContext) ComparativeExpression() IComparativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparativeExpressionContext)
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TuringParserNOT, 0)
}

func (s *NotExpressionContext) NotExpression() INotExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotExpressionContext)
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (p *TuringParser) NotExpression() (localctx INotExpressionContext) {
	localctx = NewNotExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TuringParserRULE_notExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(381)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.comparativeExpression(0)
		}

	case TuringParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.Match(TuringParserNOT)
		}
		{
			p.SetState(380)
			p.NotExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) NotExpression() INotExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(TuringParserAND, 0)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (p *TuringParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *TuringParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, TuringParserRULE_andExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.NotExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_andExpression)
			p.SetState(386)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(387)
				p.Match(TuringParserAND)
			}
			{
				p.SetState(388)
				p.NotExpression()
			}

		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *OrExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(TuringParserOR, 0)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (p *TuringParser) OrExpression() (localctx IOrExpressionContext) {
	return p.orExpression(0)
}

func (p *TuringParser) orExpression(_p int) (localctx IOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, TuringParserRULE_orExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_orExpression)
			p.SetState(397)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(398)
				p.Match(TuringParserOR)
			}
			{
				p.SetState(399)
				p.andExpression(0)
			}

		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IImpliesExpressionContext is an interface to support dynamic dispatch.
type IImpliesExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImpliesExpressionContext differentiates from other interfaces.
	IsImpliesExpressionContext()
}

type ImpliesExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpliesExpressionContext() *ImpliesExpressionContext {
	var p = new(ImpliesExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_impliesExpression
	return p
}

func (*ImpliesExpressionContext) IsImpliesExpressionContext() {}

func NewImpliesExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpliesExpressionContext {
	var p = new(ImpliesExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_impliesExpression

	return p
}

func (s *ImpliesExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpliesExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *ImpliesExpressionContext) ImpliesExpression() IImpliesExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpliesExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpliesExpressionContext)
}

func (s *ImpliesExpressionContext) IMPLIES() antlr.TerminalNode {
	return s.GetToken(TuringParserIMPLIES, 0)
}

func (s *ImpliesExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpliesExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpliesExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterImpliesExpression(s)
	}
}

func (s *ImpliesExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitImpliesExpression(s)
	}
}

func (p *TuringParser) ImpliesExpression() (localctx IImpliesExpressionContext) {
	return p.impliesExpression(0)
}

func (p *TuringParser) impliesExpression(_p int) (localctx IImpliesExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewImpliesExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IImpliesExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, TuringParserRULE_impliesExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.orExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImpliesExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_impliesExpression)
			p.SetState(408)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(409)
				p.Match(TuringParserIMPLIES)
			}
			{
				p.SetState(410)
				p.orExpression(0)
			}

		}
		p.SetState(415)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ImpliesExpression() IImpliesExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpliesExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpliesExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TuringParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, TuringParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.impliesExpression(0)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpressionListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *TuringParser) ExpressionList() (localctx IExpressionListContext) {
	return p.expressionList(0)
}

func (p *TuringParser) expressionList(_p int) (localctx IExpressionListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, TuringParserRULE_expressionList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expressionList)
			p.SetState(421)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(422)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(423)
				p.Expression()
			}

		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeDeclaration() ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *DeclarationContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *DeclarationContext) ArrayDeclaration() IArrayDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayDeclarationContext)
}

func (s *DeclarationContext) SubprogramDeclaration() ISubprogramDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubprogramDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubprogramDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *TuringParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, TuringParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(433)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)
			p.TypeDeclaration()
		}

	case TuringParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(430)
			p.VariableDeclaration()
		}

	case TuringParserARRAY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(431)
			p.ArrayDeclaration()
		}

	case TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(432)
			p.SubprogramDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementsContext) AssignmentStatement() IAssignmentStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementsContext) PutStatement() IPutStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPutStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPutStatementContext)
}

func (s *StatementsContext) EXIT() antlr.TerminalNode {
	return s.GetToken(TuringParserEXIT, 0)
}

func (s *StatementsContext) BeginStatement() IBeginStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeginStatementContext)
}

func (s *StatementsContext) RETURN() antlr.TerminalNode {
	return s.GetToken(TuringParserRETURN, 0)
}

func (s *StatementsContext) ResultStatement() IResultStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResultStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResultStatementContext)
}

func (s *StatementsContext) NewStatement() INewStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewStatementContext)
}

func (s *StatementsContext) FreeStatement() IFreeStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFreeStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFreeStatementContext)
}

func (s *StatementsContext) ForkStatement() IForkStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForkStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForkStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *TuringParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, TuringParserRULE_statements)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(436)
			p.AssignmentStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(437)
			p.PutStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(438)
			p.Match(TuringParserEXIT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(439)
			p.BeginStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(440)
			p.Match(TuringParserRETURN)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(441)
			p.ResultStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(442)
			p.NewStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(443)
			p.FreeStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(444)
			p.ForkStatement()
		}

	}

	return localctx
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_assignmentStatement
	return p
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) PostfixExpression() IPostfixExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(TuringParserASSIGNMENT, 0)
}

func (s *AssignmentStatementContext) PLUSEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserPLUSEQUALS, 0)
}

func (s *AssignmentStatementContext) MINUSEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserMINUSEQUALS, 0)
}

func (s *AssignmentStatementContext) MULTIPLYEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserMULTIPLYEQUALS, 0)
}

func (s *AssignmentStatementContext) DIVIDEEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVIDEEQUALS, 0)
}

func (s *AssignmentStatementContext) DIVEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVEQUALS, 0)
}

func (s *AssignmentStatementContext) SHLEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserSHLEQUALS, 0)
}

func (s *AssignmentStatementContext) SHREQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserSHREQUALS, 0)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (p *TuringParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, TuringParserRULE_assignmentStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.postfixExpression(0)
	}
	{
		p.SetState(448)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(TuringParserASSIGNMENT-69))|(1<<(TuringParserPLUSEQUALS-69))|(1<<(TuringParserMINUSEQUALS-69))|(1<<(TuringParserMULTIPLYEQUALS-69))|(1<<(TuringParserDIVIDEEQUALS-69))|(1<<(TuringParserDIVEQUALS-69))|(1<<(TuringParserSHREQUALS-69)))) != 0) || _la == TuringParserSHLEQUALS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(449)
		p.Expression()
	}

	return localctx
}

// IPutStatementContext is an interface to support dynamic dispatch.
type IPutStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPutStatementContext differentiates from other interfaces.
	IsPutStatementContext()
}

type PutStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPutStatementContext() *PutStatementContext {
	var p = new(PutStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_putStatement
	return p
}

func (*PutStatementContext) IsPutStatementContext() {}

func NewPutStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutStatementContext {
	var p = new(PutStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_putStatement

	return p
}

func (s *PutStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PutStatementContext) PUT() antlr.TerminalNode {
	return s.GetToken(TuringParserPUT, 0)
}

func (s *PutStatementContext) PutItemList() IPutItemListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPutItemListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPutItemListContext)
}

func (s *PutStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PutStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterPutStatement(s)
	}
}

func (s *PutStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitPutStatement(s)
	}
}

func (p *TuringParser) PutStatement() (localctx IPutStatementContext) {
	localctx = NewPutStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, TuringParserRULE_putStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(TuringParserPUT)
	}
	{
		p.SetState(452)
		p.putItemList(0)
	}

	return localctx
}

// IPutItemContext is an interface to support dynamic dispatch.
type IPutItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPutItemContext differentiates from other interfaces.
	IsPutItemContext()
}

type PutItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPutItemContext() *PutItemContext {
	var p = new(PutItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_putItem
	return p
}

func (*PutItemContext) IsPutItemContext() {}

func NewPutItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutItemContext {
	var p = new(PutItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_putItem

	return p
}

func (s *PutItemContext) GetParser() antlr.Parser { return s.parser }

func (s *PutItemContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *PutItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PutItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterPutItem(s)
	}
}

func (s *PutItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitPutItem(s)
	}
}

func (p *TuringParser) PutItem() (localctx IPutItemContext) {
	localctx = NewPutItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, TuringParserRULE_putItem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Statements()
	}

	return localctx
}

// IPutItemListContext is an interface to support dynamic dispatch.
type IPutItemListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPutItemListContext differentiates from other interfaces.
	IsPutItemListContext()
}

type PutItemListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPutItemListContext() *PutItemListContext {
	var p = new(PutItemListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_putItemList
	return p
}

func (*PutItemListContext) IsPutItemListContext() {}

func NewPutItemListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutItemListContext {
	var p = new(PutItemListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_putItemList

	return p
}

func (s *PutItemListContext) GetParser() antlr.Parser { return s.parser }

func (s *PutItemListContext) PutItem() IPutItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPutItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPutItemContext)
}

func (s *PutItemListContext) PutItemList() IPutItemListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPutItemListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPutItemListContext)
}

func (s *PutItemListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *PutItemListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutItemListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PutItemListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterPutItemList(s)
	}
}

func (s *PutItemListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitPutItemList(s)
	}
}

func (p *TuringParser) PutItemList() (localctx IPutItemListContext) {
	return p.putItemList(0)
}

func (p *TuringParser) putItemList(_p int) (localctx IPutItemListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPutItemListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPutItemListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 82
	p.EnterRecursionRule(localctx, 82, TuringParserRULE_putItemList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		p.PutItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPutItemListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_putItemList)
			p.SetState(459)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(460)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(461)
				p.PutItem()
			}

		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IBeginStatementContext is an interface to support dynamic dispatch.
type IBeginStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeginStatementContext differentiates from other interfaces.
	IsBeginStatementContext()
}

type BeginStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeginStatementContext() *BeginStatementContext {
	var p = new(BeginStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_beginStatement
	return p
}

func (*BeginStatementContext) IsBeginStatementContext() {}

func NewBeginStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginStatementContext {
	var p = new(BeginStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_beginStatement

	return p
}

func (s *BeginStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginStatementContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(TuringParserBEGIN, 0)
}

func (s *BeginStatementContext) StatementOrDeclaration() IStatementOrDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementOrDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementOrDeclarationContext)
}

func (s *BeginStatementContext) END() antlr.TerminalNode {
	return s.GetToken(TuringParserEND, 0)
}

func (s *BeginStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterBeginStatement(s)
	}
}

func (s *BeginStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitBeginStatement(s)
	}
}

func (p *TuringParser) BeginStatement() (localctx IBeginStatementContext) {
	localctx = NewBeginStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, TuringParserRULE_beginStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(467)
		p.Match(TuringParserBEGIN)
	}
	{
		p.SetState(468)
		p.StatementOrDeclaration()
	}
	{
		p.SetState(469)
		p.Match(TuringParserEND)
	}

	return localctx
}

// IResultStatementContext is an interface to support dynamic dispatch.
type IResultStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResultStatementContext differentiates from other interfaces.
	IsResultStatementContext()
}

type ResultStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultStatementContext() *ResultStatementContext {
	var p = new(ResultStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_resultStatement
	return p
}

func (*ResultStatementContext) IsResultStatementContext() {}

func NewResultStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultStatementContext {
	var p = new(ResultStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_resultStatement

	return p
}

func (s *ResultStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultStatementContext) RESULT() antlr.TerminalNode {
	return s.GetToken(TuringParserRESULT, 0)
}

func (s *ResultStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ResultStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterResultStatement(s)
	}
}

func (s *ResultStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitResultStatement(s)
	}
}

func (p *TuringParser) ResultStatement() (localctx IResultStatementContext) {
	localctx = NewResultStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, TuringParserRULE_resultStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Match(TuringParserRESULT)
	}
	{
		p.SetState(472)
		p.Expression()
	}

	return localctx
}

// INewStatementContext is an interface to support dynamic dispatch.
type INewStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewStatementContext differentiates from other interfaces.
	IsNewStatementContext()
}

type NewStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewStatementContext() *NewStatementContext {
	var p = new(NewStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_newStatement
	return p
}

func (*NewStatementContext) IsNewStatementContext() {}

func NewNewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewStatementContext {
	var p = new(NewStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_newStatement

	return p
}

func (s *NewStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *NewStatementContext) NEW() antlr.TerminalNode {
	return s.GetToken(TuringParserNEW, 0)
}

func (s *NewStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TuringParserIDENTIFIER)
}

func (s *NewStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, i)
}

func (s *NewStatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *NewStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterNewStatement(s)
	}
}

func (s *NewStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitNewStatement(s)
	}
}

func (p *TuringParser) NewStatement() (localctx INewStatementContext) {
	localctx = NewNewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, TuringParserRULE_newStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(474)
		p.Match(TuringParserNEW)
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(475)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(476)
			p.Match(TuringParserCOMMA)
		}

	}
	{
		p.SetState(479)
		p.Match(TuringParserIDENTIFIER)
	}

	return localctx
}

// IFreeStatementContext is an interface to support dynamic dispatch.
type IFreeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFreeStatementContext differentiates from other interfaces.
	IsFreeStatementContext()
}

type FreeStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFreeStatementContext() *FreeStatementContext {
	var p = new(FreeStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_freeStatement
	return p
}

func (*FreeStatementContext) IsFreeStatementContext() {}

func NewFreeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FreeStatementContext {
	var p = new(FreeStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_freeStatement

	return p
}

func (s *FreeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FreeStatementContext) FREE() antlr.TerminalNode {
	return s.GetToken(TuringParserFREE, 0)
}

func (s *FreeStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TuringParserIDENTIFIER)
}

func (s *FreeStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, i)
}

func (s *FreeStatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *FreeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FreeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FreeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterFreeStatement(s)
	}
}

func (s *FreeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitFreeStatement(s)
	}
}

func (p *TuringParser) FreeStatement() (localctx IFreeStatementContext) {
	localctx = NewFreeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, TuringParserRULE_freeStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)
		p.Match(TuringParserFREE)
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(482)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(483)
			p.Match(TuringParserCOMMA)
		}

	}
	{
		p.SetState(486)
		p.Match(TuringParserIDENTIFIER)
	}

	return localctx
}

// IForkStatementContext is an interface to support dynamic dispatch.
type IForkStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForkStatementContext differentiates from other interfaces.
	IsForkStatementContext()
}

type ForkStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForkStatementContext() *ForkStatementContext {
	var p = new(ForkStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_forkStatement
	return p
}

func (*ForkStatementContext) IsForkStatementContext() {}

func NewForkStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForkStatementContext {
	var p = new(ForkStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_forkStatement

	return p
}

func (s *ForkStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForkStatementContext) FORK() antlr.TerminalNode {
	return s.GetToken(TuringParserFORK, 0)
}

func (s *ForkStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForkStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForkStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForkStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterForkStatement(s)
	}
}

func (s *ForkStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitForkStatement(s)
	}
}

func (p *TuringParser) ForkStatement() (localctx IForkStatementContext) {
	localctx = NewForkStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, TuringParserRULE_forkStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)
		p.Match(TuringParserFORK)
	}
	{
		p.SetState(489)
		p.Expression()
	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(TuringParserTYPE, 0)
}

func (s *TypeDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *TypeDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(TuringParserCOLON, 0)
}

func (s *TypeDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (p *TuringParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, TuringParserRULE_typeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Match(TuringParserTYPE)
	}
	{
		p.SetState(492)
		p.Match(TuringParserIDENTIFIER)
	}
	{
		p.SetState(493)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(494)
		p.TypeSpec()
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) INT() antlr.TerminalNode {
	return s.GetToken(TuringParserINT, 0)
}

func (s *TypeSpecContext) REAL() antlr.TerminalNode {
	return s.GetToken(TuringParserREAL, 0)
}

func (s *TypeSpecContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(TuringParserBOOLEAN, 0)
}

func (s *TypeSpecContext) NAT() antlr.TerminalNode {
	return s.GetToken(TuringParserNAT, 0)
}

func (s *TypeSpecContext) INTN() antlr.TerminalNode {
	return s.GetToken(TuringParserINTN, 0)
}

func (s *TypeSpecContext) NATN() antlr.TerminalNode {
	return s.GetToken(TuringParserNATN, 0)
}

func (s *TypeSpecContext) REALN() antlr.TerminalNode {
	return s.GetToken(TuringParserREALN, 0)
}

func (s *TypeSpecContext) CHAR() antlr.TerminalNode {
	return s.GetToken(TuringParserCHAR, 0)
}

func (s *TypeSpecContext) IndexType() IIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeContext)
}

func (s *TypeSpecContext) StringType() IStringTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringTypeContext)
}

func (s *TypeSpecContext) RecordType() IRecordTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecordTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecordTypeContext)
}

func (s *TypeSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (p *TuringParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, TuringParserRULE_typeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(508)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.Match(TuringParserINT)
		}

	case TuringParserREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(497)
			p.Match(TuringParserREAL)
		}

	case TuringParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(498)
			p.Match(TuringParserBOOLEAN)
		}

	case TuringParserNAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(499)
			p.Match(TuringParserNAT)
		}

	case TuringParserINTN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(500)
			p.Match(TuringParserINTN)
		}

	case TuringParserNATN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(501)
			p.Match(TuringParserNATN)
		}

	case TuringParserREALN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(502)
			p.Match(TuringParserREALN)
		}

	case TuringParserCHAR:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(503)
			p.Match(TuringParserCHAR)
		}

	case TuringParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(504)
			p.IndexType()
		}

	case TuringParserSTRING:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(505)
			p.StringType()
		}

	case TuringParserRECORD:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(506)
			p.RecordType()
		}

	case TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(507)
			p.Match(TuringParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndexTypeContext is an interface to support dynamic dispatch.
type IIndexTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexTypeContext differentiates from other interfaces.
	IsIndexTypeContext()
}

type IndexTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexTypeContext() *IndexTypeContext {
	var p = new(IndexTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_indexType
	return p
}

func (*IndexTypeContext) IsIndexTypeContext() {}

func NewIndexTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTypeContext {
	var p = new(IndexTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_indexType

	return p
}

func (s *IndexTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexTypeContext) AllINTEGER_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(TuringParserINTEGER_LITERAL)
}

func (s *IndexTypeContext) INTEGER_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserINTEGER_LITERAL, i)
}

func (s *IndexTypeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(TuringParserRANGE, 0)
}

func (s *IndexTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterIndexType(s)
	}
}

func (s *IndexTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitIndexType(s)
	}
}

func (p *TuringParser) IndexType() (localctx IIndexTypeContext) {
	localctx = NewIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, TuringParserRULE_indexType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Match(TuringParserINTEGER_LITERAL)
	}
	{
		p.SetState(511)
		p.Match(TuringParserRANGE)
	}
	{
		p.SetState(512)
		p.Match(TuringParserINTEGER_LITERAL)
	}

	return localctx
}

// IIndexTypeListContext is an interface to support dynamic dispatch.
type IIndexTypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexTypeListContext differentiates from other interfaces.
	IsIndexTypeListContext()
}

type IndexTypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexTypeListContext() *IndexTypeListContext {
	var p = new(IndexTypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_indexTypeList
	return p
}

func (*IndexTypeListContext) IsIndexTypeListContext() {}

func NewIndexTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTypeListContext {
	var p = new(IndexTypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_indexTypeList

	return p
}

func (s *IndexTypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexTypeListContext) IndexType() IIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeContext)
}

func (s *IndexTypeListContext) IndexTypeList() IIndexTypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeListContext)
}

func (s *IndexTypeListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *IndexTypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterIndexTypeList(s)
	}
}

func (s *IndexTypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitIndexTypeList(s)
	}
}

func (p *TuringParser) IndexTypeList() (localctx IIndexTypeListContext) {
	return p.indexTypeList(0)
}

func (p *TuringParser) indexTypeList(_p int) (localctx IIndexTypeListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIndexTypeListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIndexTypeListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 100
	p.EnterRecursionRule(localctx, 100, TuringParserRULE_indexTypeList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)
		p.IndexType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIndexTypeListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_indexTypeList)
			p.SetState(517)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(518)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(519)
				p.IndexType()
			}

		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	return localctx
}

// IStringTypeContext is an interface to support dynamic dispatch.
type IStringTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringTypeContext differentiates from other interfaces.
	IsStringTypeContext()
}

type StringTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringTypeContext() *StringTypeContext {
	var p = new(StringTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_stringType
	return p
}

func (*StringTypeContext) IsStringTypeContext() {}

func NewStringTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTypeContext {
	var p = new(StringTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_stringType

	return p
}

func (s *StringTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(TuringParserSTRING, 0)
}

func (s *StringTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *StringTypeContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(TuringParserINTEGER_LITERAL, 0)
}

func (s *StringTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (p *TuringParser) StringType() (localctx IStringTypeContext) {
	localctx = NewStringTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, TuringParserRULE_stringType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(525)
		p.Match(TuringParserSTRING)
	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(526)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(527)
			p.Match(TuringParserINTEGER_LITERAL)
		}
		{
			p.SetState(528)
			p.Match(TuringParserR_PAREN)
		}

	}

	return localctx
}

// IRecordTypeContext is an interface to support dynamic dispatch.
type IRecordTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecordTypeContext differentiates from other interfaces.
	IsRecordTypeContext()
}

type RecordTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordTypeContext() *RecordTypeContext {
	var p = new(RecordTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_recordType
	return p
}

func (*RecordTypeContext) IsRecordTypeContext() {}

func NewRecordTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordTypeContext {
	var p = new(RecordTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_recordType

	return p
}

func (s *RecordTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordTypeContext) AllRECORD() []antlr.TerminalNode {
	return s.GetTokens(TuringParserRECORD)
}

func (s *RecordTypeContext) RECORD(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserRECORD, i)
}

func (s *RecordTypeContext) END() antlr.TerminalNode {
	return s.GetToken(TuringParserEND, 0)
}

func (s *RecordTypeContext) AllRecordField() []IRecordFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRecordFieldContext)(nil)).Elem())
	var tst = make([]IRecordFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRecordFieldContext)
		}
	}

	return tst
}

func (s *RecordTypeContext) RecordField(i int) IRecordFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecordFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRecordFieldContext)
}

func (s *RecordTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterRecordType(s)
	}
}

func (s *RecordTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitRecordType(s)
	}
}

func (p *TuringParser) RecordType() (localctx IRecordTypeContext) {
	localctx = NewRecordTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, TuringParserRULE_recordType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
		p.Match(TuringParserRECORD)
	}
	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TuringParserIDENTIFIER {
		{
			p.SetState(532)
			p.RecordField()
		}

		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(537)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(538)
		p.Match(TuringParserRECORD)
	}

	return localctx
}

// IRecordFieldContext is an interface to support dynamic dispatch.
type IRecordFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecordFieldContext differentiates from other interfaces.
	IsRecordFieldContext()
}

type RecordFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordFieldContext() *RecordFieldContext {
	var p = new(RecordFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_recordField
	return p
}

func (*RecordFieldContext) IsRecordFieldContext() {}

func NewRecordFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordFieldContext {
	var p = new(RecordFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_recordField

	return p
}

func (s *RecordFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordFieldContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *RecordFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(TuringParserCOLON, 0)
}

func (s *RecordFieldContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *RecordFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterRecordField(s)
	}
}

func (s *RecordFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitRecordField(s)
	}
}

func (p *TuringParser) RecordField() (localctx IRecordFieldContext) {
	localctx = NewRecordFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, TuringParserRULE_recordField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		p.identifierList(0)
	}
	{
		p.SetState(541)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(542)
		p.TypeSpec()
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(TuringParserVAR, 0)
}

func (s *VariableDeclarationContext) VariableIdentifier() IVariableIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableIdentifierContext)
}

func (s *VariableDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(TuringParserCOLON, 0)
}

func (s *VariableDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *VariableDeclarationContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(TuringParserASSIGNMENT, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *TuringParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, TuringParserRULE_variableDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(544)
			p.Match(TuringParserVAR)
		}
		{
			p.SetState(545)
			p.VariableIdentifier()
		}
		{
			p.SetState(546)
			p.Match(TuringParserCOLON)
		}
		{
			p.SetState(547)
			p.TypeSpec()
		}
		p.SetState(550)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserASSIGNMENT {
			{
				p.SetState(548)
				p.Match(TuringParserASSIGNMENT)
			}
			{
				p.SetState(549)
				p.Expression()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.Match(TuringParserVAR)
		}
		{
			p.SetState(553)
			p.VariableIdentifier()
		}
		{
			p.SetState(554)
			p.Match(TuringParserASSIGNMENT)
		}
		{
			p.SetState(555)
			p.Expression()
		}

	}

	return localctx
}

// IVariableIdentifierContext is an interface to support dynamic dispatch.
type IVariableIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableIdentifierContext differentiates from other interfaces.
	IsVariableIdentifierContext()
}

type VariableIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableIdentifierContext() *VariableIdentifierContext {
	var p = new(VariableIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_variableIdentifier
	return p
}

func (*VariableIdentifierContext) IsVariableIdentifierContext() {}

func NewVariableIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableIdentifierContext {
	var p = new(VariableIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_variableIdentifier

	return p
}

func (s *VariableIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *VariableIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterVariableIdentifier(s)
	}
}

func (s *VariableIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitVariableIdentifier(s)
	}
}

func (p *TuringParser) VariableIdentifier() (localctx IVariableIdentifierContext) {
	localctx = NewVariableIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, TuringParserRULE_variableIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(559)
		p.Match(TuringParserIDENTIFIER)
	}

	return localctx
}

// IArrayDeclarationContext is an interface to support dynamic dispatch.
type IArrayDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayDeclarationContext differentiates from other interfaces.
	IsArrayDeclarationContext()
}

type ArrayDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclarationContext() *ArrayDeclarationContext {
	var p = new(ArrayDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_arrayDeclaration
	return p
}

func (*ArrayDeclarationContext) IsArrayDeclarationContext() {}

func NewArrayDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclarationContext {
	var p = new(ArrayDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_arrayDeclaration

	return p
}

func (s *ArrayDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclarationContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(TuringParserARRAY, 0)
}

func (s *ArrayDeclarationContext) IndexType() IIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeContext)
}

func (s *ArrayDeclarationContext) OF() antlr.TerminalNode {
	return s.GetToken(TuringParserOF, 0)
}

func (s *ArrayDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ArrayDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *ArrayDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterArrayDeclaration(s)
	}
}

func (s *ArrayDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitArrayDeclaration(s)
	}
}

func (p *TuringParser) ArrayDeclaration() (localctx IArrayDeclarationContext) {
	localctx = NewArrayDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, TuringParserRULE_arrayDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(561)
			p.Match(TuringParserARRAY)
		}
		{
			p.SetState(562)
			p.IndexType()
		}
		{
			p.SetState(563)
			p.Match(TuringParserOF)
		}
		{
			p.SetState(564)
			p.TypeSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(566)
			p.Match(TuringParserARRAY)
		}
		{
			p.SetState(567)
			p.IndexType()
		}
		{
			p.SetState(568)
			p.Match(TuringParserCOMMA)
		}

	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *IdentifierListContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *IdentifierListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (p *TuringParser) IdentifierList() (localctx IIdentifierListContext) {
	return p.identifierList(0)
}

func (p *TuringParser) identifierList(_p int) (localctx IIdentifierListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIdentifierListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 114
	p.EnterRecursionRule(localctx, 114, TuringParserRULE_identifierList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Match(TuringParserIDENTIFIER)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIdentifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_identifierList)
			p.SetState(575)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(576)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(577)
				p.Match(TuringParserIDENTIFIER)
			}

		}
		p.SetState(582)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(TuringParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(TuringParserINTEGER_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *TuringParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, TuringParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(583)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TuringParserSTRING_LITERAL || _la == TuringParserINTEGER_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) BLOCK_COMMENT() antlr.TerminalNode {
	return s.GetToken(TuringParserBLOCK_COMMENT, 0)
}

func (s *CommentContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(TuringParserLINE_COMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *TuringParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, TuringParserRULE_comment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TuringParserBLOCK_COMMENT || _la == TuringParserLINE_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *TuringParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ExportListContext = nil
		if localctx != nil {
			t = localctx.(*ExportListContext)
		}
		return p.ExportList_Sempred(t, predIndex)

	case 17:
		var t *ParamDeclarationListContext = nil
		if localctx != nil {
			t = localctx.(*ParamDeclarationListContext)
		}
		return p.ParamDeclarationList_Sempred(t, predIndex)

	case 21:
		var t *ArgumentListContext = nil
		if localctx != nil {
			t = localctx.(*ArgumentListContext)
		}
		return p.ArgumentList_Sempred(t, predIndex)

	case 24:
		var t *PostfixExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PostfixExpressionContext)
		}
		return p.PostfixExpression_Sempred(t, predIndex)

	case 26:
		var t *MultiplicativeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeExpressionContext)
		}
		return p.MultiplicativeExpression_Sempred(t, predIndex)

	case 28:
		var t *AdditiveExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveExpressionContext)
		}
		return p.AdditiveExpression_Sempred(t, predIndex)

	case 29:
		var t *ComparativeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ComparativeExpressionContext)
		}
		return p.ComparativeExpression_Sempred(t, predIndex)

	case 31:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 32:
		var t *OrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*OrExpressionContext)
		}
		return p.OrExpression_Sempred(t, predIndex)

	case 33:
		var t *ImpliesExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ImpliesExpressionContext)
		}
		return p.ImpliesExpression_Sempred(t, predIndex)

	case 35:
		var t *ExpressionListContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionListContext)
		}
		return p.ExpressionList_Sempred(t, predIndex)

	case 41:
		var t *PutItemListContext = nil
		if localctx != nil {
			t = localctx.(*PutItemListContext)
		}
		return p.PutItemList_Sempred(t, predIndex)

	case 50:
		var t *IndexTypeListContext = nil
		if localctx != nil {
			t = localctx.(*IndexTypeListContext)
		}
		return p.IndexTypeList_Sempred(t, predIndex)

	case 57:
		var t *IdentifierListContext = nil
		if localctx != nil {
			t = localctx.(*IdentifierListContext)
		}
		return p.IdentifierList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TuringParser) ExportList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) ParamDeclarationList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) ArgumentList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) PostfixExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) MultiplicativeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) AdditiveExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) ComparativeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) OrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) ImpliesExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) ExpressionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) PutItemList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) IndexTypeList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) IdentifierList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
