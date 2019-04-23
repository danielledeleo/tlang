// Code generated from tlang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // tlang

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 120, 496,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 6, 2, 106, 10, 2, 13, 2, 14,
	2, 107, 3, 3, 3, 3, 3, 3, 5, 3, 113, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 119, 10, 4, 3, 4, 5, 4, 122, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 128,
	10, 5, 3, 5, 5, 5, 131, 10, 5, 3, 5, 5, 5, 134, 10, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 5, 6, 140, 10, 6, 3, 6, 3, 6, 5, 6, 144, 10, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 6, 9, 156, 10, 9, 13, 9, 14,
	9, 157, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 167, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 176, 10, 11, 3,
	12, 5, 12, 179, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 7, 13, 189, 10, 13, 12, 13, 14, 13, 192, 11, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 201, 10, 14, 3, 15, 3, 15, 5, 15,
	205, 10, 15, 3, 15, 3, 15, 3, 15, 5, 15, 210, 10, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 216, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 222,
	10, 16, 3, 17, 3, 17, 3, 18, 5, 18, 227, 10, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 5, 18, 234, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	7, 19, 242, 10, 19, 12, 19, 14, 19, 245, 11, 19, 3, 20, 6, 20, 248, 10,
	20, 13, 20, 14, 20, 249, 3, 21, 3, 21, 5, 21, 254, 10, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 262, 10, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 272, 10, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 304, 10, 23, 3, 23, 7,
	23, 307, 10, 23, 12, 23, 14, 23, 310, 11, 23, 3, 24, 3, 24, 3, 24, 7, 24,
	315, 10, 24, 12, 24, 14, 24, 318, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 325, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 5, 26, 336, 10, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 351, 10, 30,
	12, 30, 14, 30, 354, 11, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 32, 3, 33, 3, 33, 3, 33, 5, 33, 366, 10, 33, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 34, 5, 34, 373, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 396, 10, 39, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 408, 10, 41, 12,
	41, 14, 41, 411, 11, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 418,
	10, 42, 3, 43, 3, 43, 6, 43, 422, 10, 43, 13, 43, 14, 43, 423, 3, 43, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 5, 45, 439, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 446,
	10, 45, 3, 46, 3, 46, 3, 46, 7, 46, 451, 10, 46, 12, 46, 14, 46, 454, 11,
	46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 5, 48, 467, 10, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7,
	49, 475, 10, 49, 12, 49, 14, 49, 478, 11, 49, 3, 50, 3, 50, 3, 50, 5, 50,
	483, 10, 50, 3, 51, 3, 51, 5, 51, 487, 10, 51, 3, 51, 3, 51, 3, 51, 5,
	51, 492, 10, 51, 3, 52, 3, 52, 3, 52, 2, 8, 24, 36, 44, 58, 80, 96, 53,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
	76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 2, 14, 4, 2,
	18, 18, 21, 21, 3, 2, 70, 72, 4, 2, 3, 5, 17, 17, 4, 2, 81, 81, 84, 85,
	4, 2, 86, 90, 102, 103, 4, 2, 84, 85, 104, 104, 5, 2, 6, 6, 92, 97, 101,
	101, 5, 2, 83, 83, 105, 109, 114, 115, 4, 2, 78, 78, 82, 82, 4, 2, 49,
	51, 58, 69, 3, 2, 84, 85, 3, 2, 119, 120, 2, 523, 2, 105, 3, 2, 2, 2, 4,
	112, 3, 2, 2, 2, 6, 114, 3, 2, 2, 2, 8, 123, 3, 2, 2, 2, 10, 139, 3, 2,
	2, 2, 12, 145, 3, 2, 2, 2, 14, 150, 3, 2, 2, 2, 16, 152, 3, 2, 2, 2, 18,
	166, 3, 2, 2, 2, 20, 175, 3, 2, 2, 2, 22, 178, 3, 2, 2, 2, 24, 182, 3,
	2, 2, 2, 26, 200, 3, 2, 2, 2, 28, 215, 3, 2, 2, 2, 30, 221, 3, 2, 2, 2,
	32, 223, 3, 2, 2, 2, 34, 233, 3, 2, 2, 2, 36, 235, 3, 2, 2, 2, 38, 247,
	3, 2, 2, 2, 40, 253, 3, 2, 2, 2, 42, 261, 3, 2, 2, 2, 44, 271, 3, 2, 2,
	2, 46, 311, 3, 2, 2, 2, 48, 324, 3, 2, 2, 2, 50, 335, 3, 2, 2, 2, 52, 337,
	3, 2, 2, 2, 54, 339, 3, 2, 2, 2, 56, 342, 3, 2, 2, 2, 58, 344, 3, 2, 2,
	2, 60, 355, 3, 2, 2, 2, 62, 359, 3, 2, 2, 2, 64, 362, 3, 2, 2, 2, 66, 369,
	3, 2, 2, 2, 68, 376, 3, 2, 2, 2, 70, 379, 3, 2, 2, 2, 72, 384, 3, 2, 2,
	2, 74, 386, 3, 2, 2, 2, 76, 395, 3, 2, 2, 2, 78, 397, 3, 2, 2, 2, 80, 401,
	3, 2, 2, 2, 82, 412, 3, 2, 2, 2, 84, 419, 3, 2, 2, 2, 86, 428, 3, 2, 2,
	2, 88, 445, 3, 2, 2, 2, 90, 447, 3, 2, 2, 2, 92, 455, 3, 2, 2, 2, 94, 466,
	3, 2, 2, 2, 96, 468, 3, 2, 2, 2, 98, 482, 3, 2, 2, 2, 100, 491, 3, 2, 2,
	2, 102, 493, 3, 2, 2, 2, 104, 106, 5, 4, 3, 2, 105, 104, 3, 2, 2, 2, 106,
	107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 3, 3,
	2, 2, 2, 109, 113, 5, 40, 21, 2, 110, 113, 5, 16, 9, 2, 111, 113, 5, 102,
	52, 2, 112, 109, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 111, 3, 2, 2, 2,
	113, 5, 3, 2, 2, 2, 114, 115, 9, 2, 2, 2, 115, 121, 7, 117, 2, 2, 116,
	118, 7, 76, 2, 2, 117, 119, 5, 36, 19, 2, 118, 117, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 7, 77, 2, 2, 121, 116, 3, 2,
	2, 2, 121, 122, 3, 2, 2, 2, 122, 7, 3, 2, 2, 2, 123, 124, 7, 19, 2, 2,
	124, 130, 7, 117, 2, 2, 125, 127, 7, 76, 2, 2, 126, 128, 5, 36, 19, 2,
	127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	131, 7, 77, 2, 2, 130, 125, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 133,
	3, 2, 2, 2, 132, 134, 7, 117, 2, 2, 133, 132, 3, 2, 2, 2, 133, 134, 3,
	2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 7, 75, 2, 2, 136, 137, 5, 76,
	39, 2, 137, 9, 3, 2, 2, 2, 138, 140, 5, 14, 8, 2, 139, 138, 3, 2, 2, 2,
	139, 140, 3, 2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 144, 5, 8, 5, 2, 142,
	144, 5, 6, 4, 2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2, 2, 2, 144, 11, 3,
	2, 2, 2, 145, 146, 5, 10, 6, 2, 146, 147, 5, 38, 20, 2, 147, 148, 7, 13,
	2, 2, 148, 149, 7, 117, 2, 2, 149, 13, 3, 2, 2, 2, 150, 151, 9, 3, 2, 2,
	151, 15, 3, 2, 2, 2, 152, 153, 7, 20, 2, 2, 153, 155, 7, 117, 2, 2, 154,
	156, 5, 18, 10, 2, 155, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 155,
	3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 7, 13,
	2, 2, 160, 161, 7, 117, 2, 2, 161, 17, 3, 2, 2, 2, 162, 167, 5, 20, 11,
	2, 163, 167, 5, 26, 14, 2, 164, 167, 5, 28, 15, 2, 165, 167, 5, 40, 21,
	2, 166, 162, 3, 2, 2, 2, 166, 163, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166,
	165, 3, 2, 2, 2, 167, 19, 3, 2, 2, 2, 168, 169, 7, 43, 2, 2, 169, 176,
	5, 24, 13, 2, 170, 171, 7, 43, 2, 2, 171, 172, 7, 76, 2, 2, 172, 173, 5,
	24, 13, 2, 173, 174, 7, 77, 2, 2, 174, 176, 3, 2, 2, 2, 175, 168, 3, 2,
	2, 2, 175, 170, 3, 2, 2, 2, 176, 21, 3, 2, 2, 2, 177, 179, 5, 32, 17, 2,
	178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180,
	181, 7, 117, 2, 2, 181, 23, 3, 2, 2, 2, 182, 183, 8, 13, 1, 2, 183, 184,
	5, 22, 12, 2, 184, 190, 3, 2, 2, 2, 185, 186, 12, 3, 2, 2, 186, 187, 7,
	80, 2, 2, 187, 189, 5, 22, 12, 2, 188, 185, 3, 2, 2, 2, 189, 192, 3, 2,
	2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 25, 3, 2, 2, 2,
	192, 190, 3, 2, 2, 2, 193, 194, 7, 45, 2, 2, 194, 201, 5, 30, 16, 2, 195,
	196, 7, 45, 2, 2, 196, 197, 7, 76, 2, 2, 197, 198, 5, 30, 16, 2, 198, 199,
	7, 77, 2, 2, 199, 201, 3, 2, 2, 2, 200, 193, 3, 2, 2, 2, 200, 195, 3, 2,
	2, 2, 201, 27, 3, 2, 2, 2, 202, 204, 7, 46, 2, 2, 203, 205, 7, 47, 2, 2,
	204, 203, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206,
	216, 5, 30, 16, 2, 207, 209, 7, 46, 2, 2, 208, 210, 7, 47, 2, 2, 209, 208,
	3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 7, 76,
	2, 2, 212, 213, 5, 30, 16, 2, 213, 214, 7, 77, 2, 2, 214, 216, 3, 2, 2,
	2, 215, 202, 3, 2, 2, 2, 215, 207, 3, 2, 2, 2, 216, 29, 3, 2, 2, 2, 217,
	222, 7, 117, 2, 2, 218, 219, 7, 117, 2, 2, 219, 220, 7, 101, 2, 2, 220,
	222, 7, 7, 2, 2, 221, 217, 3, 2, 2, 2, 221, 218, 3, 2, 2, 2, 222, 31, 3,
	2, 2, 2, 223, 224, 9, 4, 2, 2, 224, 33, 3, 2, 2, 2, 225, 227, 7, 17, 2,
	2, 226, 225, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228,
	229, 5, 96, 49, 2, 229, 230, 7, 75, 2, 2, 230, 231, 5, 76, 39, 2, 231,
	234, 3, 2, 2, 2, 232, 234, 5, 10, 6, 2, 233, 226, 3, 2, 2, 2, 233, 232,
	3, 2, 2, 2, 234, 35, 3, 2, 2, 2, 235, 236, 8, 19, 1, 2, 236, 237, 5, 34,
	18, 2, 237, 243, 3, 2, 2, 2, 238, 239, 12, 3, 2, 2, 239, 240, 7, 80, 2,
	2, 240, 242, 5, 34, 18, 2, 241, 238, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2,
	243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 37, 3, 2, 2, 2, 245, 243,
	3, 2, 2, 2, 246, 248, 5, 40, 21, 2, 247, 246, 3, 2, 2, 2, 248, 249, 3,
	2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 39, 3, 2, 2,
	2, 251, 254, 5, 50, 26, 2, 252, 254, 5, 48, 25, 2, 253, 251, 3, 2, 2, 2,
	253, 252, 3, 2, 2, 2, 254, 41, 3, 2, 2, 2, 255, 262, 5, 98, 50, 2, 256,
	262, 7, 117, 2, 2, 257, 258, 7, 76, 2, 2, 258, 259, 5, 44, 23, 2, 259,
	260, 7, 77, 2, 2, 260, 262, 3, 2, 2, 2, 261, 255, 3, 2, 2, 2, 261, 256,
	3, 2, 2, 2, 261, 257, 3, 2, 2, 2, 262, 43, 3, 2, 2, 2, 263, 264, 8, 23,
	1, 2, 264, 272, 5, 42, 22, 2, 265, 266, 7, 74, 2, 2, 266, 272, 5, 44, 23,
	15, 267, 268, 9, 5, 2, 2, 268, 272, 5, 44, 23, 11, 269, 270, 7, 73, 2,
	2, 270, 272, 5, 44, 23, 7, 271, 263, 3, 2, 2, 2, 271, 265, 3, 2, 2, 2,
	271, 267, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272, 308, 3, 2, 2, 2, 273,
	274, 12, 13, 2, 2, 274, 275, 7, 91, 2, 2, 275, 307, 5, 44, 23, 14, 276,
	277, 12, 10, 2, 2, 277, 278, 9, 6, 2, 2, 278, 307, 5, 44, 23, 11, 279,
	280, 12, 9, 2, 2, 280, 281, 9, 7, 2, 2, 281, 307, 5, 44, 23, 10, 282, 283,
	12, 8, 2, 2, 283, 284, 9, 8, 2, 2, 284, 307, 5, 44, 23, 9, 285, 286, 12,
	6, 2, 2, 286, 287, 7, 98, 2, 2, 287, 307, 5, 44, 23, 7, 288, 289, 12, 5,
	2, 2, 289, 290, 7, 99, 2, 2, 290, 307, 5, 44, 23, 6, 291, 292, 12, 4, 2,
	2, 292, 293, 7, 100, 2, 2, 293, 307, 5, 44, 23, 5, 294, 295, 12, 3, 2,
	2, 295, 296, 9, 9, 2, 2, 296, 307, 5, 44, 23, 4, 297, 298, 12, 14, 2, 2,
	298, 299, 9, 10, 2, 2, 299, 307, 7, 117, 2, 2, 300, 301, 12, 12, 2, 2,
	301, 303, 7, 76, 2, 2, 302, 304, 5, 46, 24, 2, 303, 302, 3, 2, 2, 2, 303,
	304, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307, 7, 77, 2, 2, 306, 273,
	3, 2, 2, 2, 306, 276, 3, 2, 2, 2, 306, 279, 3, 2, 2, 2, 306, 282, 3, 2,
	2, 2, 306, 285, 3, 2, 2, 2, 306, 288, 3, 2, 2, 2, 306, 291, 3, 2, 2, 2,
	306, 294, 3, 2, 2, 2, 306, 297, 3, 2, 2, 2, 306, 300, 3, 2, 2, 2, 307,
	310, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 45, 3,
	2, 2, 2, 310, 308, 3, 2, 2, 2, 311, 316, 5, 44, 23, 2, 312, 313, 7, 80,
	2, 2, 313, 315, 5, 44, 23, 2, 314, 312, 3, 2, 2, 2, 315, 318, 3, 2, 2,
	2, 316, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 47, 3, 2, 2, 2, 318,
	316, 3, 2, 2, 2, 319, 325, 5, 70, 36, 2, 320, 325, 5, 88, 45, 2, 321, 325,
	5, 94, 48, 2, 322, 325, 5, 12, 7, 2, 323, 325, 5, 16, 9, 2, 324, 319, 3,
	2, 2, 2, 324, 320, 3, 2, 2, 2, 324, 321, 3, 2, 2, 2, 324, 322, 3, 2, 2,
	2, 324, 323, 3, 2, 2, 2, 325, 49, 3, 2, 2, 2, 326, 336, 5, 44, 23, 2, 327,
	336, 5, 54, 28, 2, 328, 336, 7, 24, 2, 2, 329, 336, 5, 60, 31, 2, 330,
	336, 7, 31, 2, 2, 331, 336, 5, 62, 32, 2, 332, 336, 5, 64, 33, 2, 333,
	336, 5, 66, 34, 2, 334, 336, 5, 68, 35, 2, 335, 326, 3, 2, 2, 2, 335, 327,
	3, 2, 2, 2, 335, 328, 3, 2, 2, 2, 335, 329, 3, 2, 2, 2, 335, 330, 3, 2,
	2, 2, 335, 331, 3, 2, 2, 2, 335, 332, 3, 2, 2, 2, 335, 333, 3, 2, 2, 2,
	335, 334, 3, 2, 2, 2, 336, 51, 3, 2, 2, 2, 337, 338, 9, 9, 2, 2, 338, 53,
	3, 2, 2, 2, 339, 340, 7, 48, 2, 2, 340, 341, 5, 58, 30, 2, 341, 55, 3,
	2, 2, 2, 342, 343, 5, 50, 26, 2, 343, 57, 3, 2, 2, 2, 344, 345, 8, 30,
	1, 2, 345, 346, 5, 56, 29, 2, 346, 352, 3, 2, 2, 2, 347, 348, 12, 3, 2,
	2, 348, 349, 7, 80, 2, 2, 349, 351, 5, 56, 29, 2, 350, 347, 3, 2, 2, 2,
	351, 354, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353,
	59, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2, 355, 356, 7, 30, 2, 2, 356, 357,
	5, 40, 21, 2, 357, 358, 7, 13, 2, 2, 358, 61, 3, 2, 2, 2, 359, 360, 7,
	32, 2, 2, 360, 361, 5, 44, 23, 2, 361, 63, 3, 2, 2, 2, 362, 365, 7, 33,
	2, 2, 363, 364, 7, 117, 2, 2, 364, 366, 7, 80, 2, 2, 365, 363, 3, 2, 2,
	2, 365, 366, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 368, 7, 117, 2, 2,
	368, 65, 3, 2, 2, 2, 369, 372, 7, 34, 2, 2, 370, 371, 7, 117, 2, 2, 371,
	373, 7, 80, 2, 2, 372, 370, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 374,
	3, 2, 2, 2, 374, 375, 7, 117, 2, 2, 375, 67, 3, 2, 2, 2, 376, 377, 7, 36,
	2, 2, 377, 378, 5, 44, 23, 2, 378, 69, 3, 2, 2, 2, 379, 380, 7, 16, 2,
	2, 380, 381, 7, 117, 2, 2, 381, 382, 7, 75, 2, 2, 382, 383, 5, 76, 39,
	2, 383, 71, 3, 2, 2, 2, 384, 385, 9, 11, 2, 2, 385, 73, 3, 2, 2, 2, 386,
	387, 7, 117, 2, 2, 387, 75, 3, 2, 2, 2, 388, 396, 5, 72, 37, 2, 389, 396,
	5, 78, 40, 2, 390, 396, 5, 82, 42, 2, 391, 396, 5, 84, 43, 2, 392, 396,
	5, 94, 48, 2, 393, 396, 5, 16, 9, 2, 394, 396, 5, 74, 38, 2, 395, 388,
	3, 2, 2, 2, 395, 389, 3, 2, 2, 2, 395, 390, 3, 2, 2, 2, 395, 391, 3, 2,
	2, 2, 395, 392, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 395, 394, 3, 2, 2, 2,
	396, 77, 3, 2, 2, 2, 397, 398, 5, 100, 51, 2, 398, 399, 7, 79, 2, 2, 399,
	400, 5, 100, 51, 2, 400, 79, 3, 2, 2, 2, 401, 402, 8, 41, 1, 2, 402, 403,
	5, 78, 40, 2, 403, 409, 3, 2, 2, 2, 404, 405, 12, 3, 2, 2, 405, 406, 7,
	80, 2, 2, 406, 408, 5, 78, 40, 2, 407, 404, 3, 2, 2, 2, 408, 411, 3, 2,
	2, 2, 409, 407, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 81, 3, 2, 2, 2,
	411, 409, 3, 2, 2, 2, 412, 417, 7, 52, 2, 2, 413, 414, 7, 76, 2, 2, 414,
	415, 5, 100, 51, 2, 415, 416, 7, 77, 2, 2, 416, 418, 3, 2, 2, 2, 417, 413,
	3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 83, 3, 2, 2, 2, 419, 421, 7, 55,
	2, 2, 420, 422, 5, 86, 44, 2, 421, 420, 3, 2, 2, 2, 422, 423, 3, 2, 2,
	2, 423, 421, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425,
	426, 7, 13, 2, 2, 426, 427, 7, 55, 2, 2, 427, 85, 3, 2, 2, 2, 428, 429,
	5, 96, 49, 2, 429, 430, 7, 75, 2, 2, 430, 431, 5, 76, 39, 2, 431, 87, 3,
	2, 2, 2, 432, 433, 7, 17, 2, 2, 433, 434, 5, 90, 46, 2, 434, 435, 7, 75,
	2, 2, 435, 438, 5, 76, 39, 2, 436, 437, 7, 83, 2, 2, 437, 439, 5, 44, 23,
	2, 438, 436, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 446, 3, 2, 2, 2, 440,
	441, 7, 17, 2, 2, 441, 442, 5, 90, 46, 2, 442, 443, 7, 83, 2, 2, 443, 444,
	5, 44, 23, 2, 444, 446, 3, 2, 2, 2, 445, 432, 3, 2, 2, 2, 445, 440, 3,
	2, 2, 2, 446, 89, 3, 2, 2, 2, 447, 452, 5, 92, 47, 2, 448, 449, 7, 80,
	2, 2, 449, 451, 5, 92, 47, 2, 450, 448, 3, 2, 2, 2, 451, 454, 3, 2, 2,
	2, 452, 450, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 91, 3, 2, 2, 2, 454,
	452, 3, 2, 2, 2, 455, 456, 7, 117, 2, 2, 456, 93, 3, 2, 2, 2, 457, 458,
	7, 53, 2, 2, 458, 459, 5, 78, 40, 2, 459, 460, 7, 14, 2, 2, 460, 461, 5,
	76, 39, 2, 461, 467, 3, 2, 2, 2, 462, 463, 7, 53, 2, 2, 463, 464, 5, 78,
	40, 2, 464, 465, 7, 80, 2, 2, 465, 467, 3, 2, 2, 2, 466, 457, 3, 2, 2,
	2, 466, 462, 3, 2, 2, 2, 467, 95, 3, 2, 2, 2, 468, 469, 8, 49, 1, 2, 469,
	470, 7, 117, 2, 2, 470, 476, 3, 2, 2, 2, 471, 472, 12, 3, 2, 2, 472, 473,
	7, 80, 2, 2, 473, 475, 7, 117, 2, 2, 474, 471, 3, 2, 2, 2, 475, 478, 3,
	2, 2, 2, 476, 474, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2, 477, 97, 3, 2, 2,
	2, 478, 476, 3, 2, 2, 2, 479, 483, 7, 7, 2, 2, 480, 483, 7, 12, 2, 2, 481,
	483, 5, 100, 51, 2, 482, 479, 3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 482, 481,
	3, 2, 2, 2, 483, 99, 3, 2, 2, 2, 484, 492, 7, 8, 2, 2, 485, 487, 9, 12,
	2, 2, 486, 485, 3, 2, 2, 2, 486, 487, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2,
	488, 492, 7, 9, 2, 2, 489, 492, 7, 11, 2, 2, 490, 492, 7, 10, 2, 2, 491,
	484, 3, 2, 2, 2, 491, 486, 3, 2, 2, 2, 491, 489, 3, 2, 2, 2, 491, 490,
	3, 2, 2, 2, 492, 101, 3, 2, 2, 2, 493, 494, 9, 13, 2, 2, 494, 103, 3, 2,
	2, 2, 49, 107, 112, 118, 121, 127, 130, 133, 139, 143, 157, 166, 175, 178,
	190, 200, 204, 209, 215, 221, 226, 233, 243, 249, 253, 261, 271, 303, 306,
	308, 316, 324, 335, 352, 365, 372, 395, 409, 417, 423, 438, 445, 452, 466,
	476, 482, 486, 491,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'unqualified'", "'pervasive'", "'opaque'", "", "", "", "", "", "",
	"", "'end'", "'of'", "'to'", "'type'", "'var'", "", "", "'class'", "'process'",
	"'for'", "'loop'", "'exit'", "'if'", "'else'", "'elsif'", "'case'", "'assert'",
	"'begin'", "'return'", "'result'", "'new'", "'free'", "'tag'", "'fork'",
	"'signal'", "'wait'", "'pause'", "'quit'", "'unchecked'", "'checked'",
	"'export'", "'import'", "'inherit'", "'implement'", "'by'", "'put'", "'int'",
	"'real'", "'boolean'", "'string'", "'array'", "'set'", "'record'", "'union'",
	"'pointer'", "'nat'", "'int1'", "'int2'", "'int4'", "'int8'", "'nat1'",
	"'nat2'", "'nat4'", "'nat8'", "'real4'", "'real8'", "'char'", "'deferred'",
	"'forward'", "'body'", "'not'", "'^'", "':'", "'('", "')'", "'.'", "'..'",
	"','", "'#'", "'->'", "':='", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'",
	"'rem'", "'**'", "'<'", "'>'", "'='", "'<='", "'>='", "'not='", "'and'",
	"'or'", "'=>'", "'in'", "'shr'", "'shl'", "'xor'", "'+='", "'-='", "'*='",
	"'/='", "'div='", "'mod='", "'and='", "'or='", "'=>='", "'shr='", "'shl='",
	"'xor='",
}
var symbolicNames = []string{
	"", "", "", "", "NOT_IN", "STRING_LITERAL", "OCTAL_LITERAL", "DECIMAL_LITERAL",
	"HEX_LITERAL", "BINARY_LITERAL", "REAL_LITERAL", "END", "OF", "TO", "TYPE",
	"VAR", "PROCEDURE", "FUNCTION", "CLASS", "PROCESS", "FOR", "LOOP", "EXIT",
	"IF", "ELSE", "ELSIF", "CASE", "ASSERT", "BEGIN", "RETURN", "RESULT", "NEW",
	"FREE", "TAG", "FORK", "SIGNAL", "WAIT", "PAUSE", "QUIT", "UNCHECKED",
	"CHECKED", "EXPORT", "IMPORT", "INHERIT", "IMPLEMENT", "BY", "PUT", "INT",
	"REAL", "BOOLEAN", "STRING", "ARRAY", "SET", "RECORD", "UNION", "POINTER",
	"NAT", "INT1", "INT2", "INT4", "INT8", "NAT1", "NAT2", "NAT4", "NAT8",
	"REAL4", "REAL8", "CHAR", "DEFERRED", "FORWARD", "BODY", "NOT", "CARET",
	"COLON", "L_PAREN", "R_PAREN", "DOT", "RANGE", "COMMA", "CHEAT", "DEREFERENCE",
	"ASSIGNMENT", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "DIV", "MOD", "REM",
	"EXP", "LESSTHAN", "GREATERTHAN", "EQUAL", "LESSTHANEQUAL", "GREATERTHANEQUAL",
	"NOTEQUAL", "AND", "OR", "IMPLIES", "IN", "SHR", "SHL", "XOR", "PLUSEQUALS",
	"MINUSEQUALS", "MULTIPLYEQUALS", "DIVIDEEQUALS", "DIVEQUALS", "MODEQUALS",
	"ANDEQUALS", "OREQUALS", "IMPLIESEQUALS", "SHREQUALS", "SHLEQUALS", "XOREQUALS",
	"IDENTIFIER", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"program", "topLevel", "procHeader", "funcHeader", "subprogramHeader",
	"subprogramDeclaration", "subprogramPrefix", "classDeclaration", "classBody",
	"exportStatement", "exportListItem", "exportList", "inheritStatement",
	"implementStatement", "idOrFileItem", "howExport", "paramDeclaration",
	"paramDeclarationList", "procBody", "statementOrDeclaration", "primaryExpression",
	"expression", "expressionList", "declaration", "statement", "assignmentOperator",
	"putStatement", "putItem", "putItemList", "beginStatement", "resultStatement",
	"newStatement", "freeStatement", "forkStatement", "typeDeclaration", "basicType",
	"referenceType", "typeSpec", "indexType", "indexTypeList", "stringType",
	"recordType", "recordField", "variableDeclaration", "variableIdentifierList",
	"variableIdentifier", "arrayDeclaration", "identifierList", "literal",
	"integer_literal", "comment",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type tlangParser struct {
	*antlr.BaseParser
}

func NewtlangParser(input antlr.TokenStream) *tlangParser {
	this := new(tlangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "tlang.g4"

	return this
}

// tlangParser tokens.
const (
	tlangParserEOF              = antlr.TokenEOF
	tlangParserT__0             = 1
	tlangParserT__1             = 2
	tlangParserT__2             = 3
	tlangParserNOT_IN           = 4
	tlangParserSTRING_LITERAL   = 5
	tlangParserOCTAL_LITERAL    = 6
	tlangParserDECIMAL_LITERAL  = 7
	tlangParserHEX_LITERAL      = 8
	tlangParserBINARY_LITERAL   = 9
	tlangParserREAL_LITERAL     = 10
	tlangParserEND              = 11
	tlangParserOF               = 12
	tlangParserTO               = 13
	tlangParserTYPE             = 14
	tlangParserVAR              = 15
	tlangParserPROCEDURE        = 16
	tlangParserFUNCTION         = 17
	tlangParserCLASS            = 18
	tlangParserPROCESS          = 19
	tlangParserFOR              = 20
	tlangParserLOOP             = 21
	tlangParserEXIT             = 22
	tlangParserIF               = 23
	tlangParserELSE             = 24
	tlangParserELSIF            = 25
	tlangParserCASE             = 26
	tlangParserASSERT           = 27
	tlangParserBEGIN            = 28
	tlangParserRETURN           = 29
	tlangParserRESULT           = 30
	tlangParserNEW              = 31
	tlangParserFREE             = 32
	tlangParserTAG              = 33
	tlangParserFORK             = 34
	tlangParserSIGNAL           = 35
	tlangParserWAIT             = 36
	tlangParserPAUSE            = 37
	tlangParserQUIT             = 38
	tlangParserUNCHECKED        = 39
	tlangParserCHECKED          = 40
	tlangParserEXPORT           = 41
	tlangParserIMPORT           = 42
	tlangParserINHERIT          = 43
	tlangParserIMPLEMENT        = 44
	tlangParserBY               = 45
	tlangParserPUT              = 46
	tlangParserINT              = 47
	tlangParserREAL             = 48
	tlangParserBOOLEAN          = 49
	tlangParserSTRING           = 50
	tlangParserARRAY            = 51
	tlangParserSET              = 52
	tlangParserRECORD           = 53
	tlangParserUNION            = 54
	tlangParserPOINTER          = 55
	tlangParserNAT              = 56
	tlangParserINT1             = 57
	tlangParserINT2             = 58
	tlangParserINT4             = 59
	tlangParserINT8             = 60
	tlangParserNAT1             = 61
	tlangParserNAT2             = 62
	tlangParserNAT4             = 63
	tlangParserNAT8             = 64
	tlangParserREAL4            = 65
	tlangParserREAL8            = 66
	tlangParserCHAR             = 67
	tlangParserDEFERRED         = 68
	tlangParserFORWARD          = 69
	tlangParserBODY             = 70
	tlangParserNOT              = 71
	tlangParserCARET            = 72
	tlangParserCOLON            = 73
	tlangParserL_PAREN          = 74
	tlangParserR_PAREN          = 75
	tlangParserDOT              = 76
	tlangParserRANGE            = 77
	tlangParserCOMMA            = 78
	tlangParserCHEAT            = 79
	tlangParserDEREFERENCE      = 80
	tlangParserASSIGNMENT       = 81
	tlangParserPLUS             = 82
	tlangParserMINUS            = 83
	tlangParserMULTIPLY         = 84
	tlangParserDIVIDE           = 85
	tlangParserDIV              = 86
	tlangParserMOD              = 87
	tlangParserREM              = 88
	tlangParserEXP              = 89
	tlangParserLESSTHAN         = 90
	tlangParserGREATERTHAN      = 91
	tlangParserEQUAL            = 92
	tlangParserLESSTHANEQUAL    = 93
	tlangParserGREATERTHANEQUAL = 94
	tlangParserNOTEQUAL         = 95
	tlangParserAND              = 96
	tlangParserOR               = 97
	tlangParserIMPLIES          = 98
	tlangParserIN               = 99
	tlangParserSHR              = 100
	tlangParserSHL              = 101
	tlangParserXOR              = 102
	tlangParserPLUSEQUALS       = 103
	tlangParserMINUSEQUALS      = 104
	tlangParserMULTIPLYEQUALS   = 105
	tlangParserDIVIDEEQUALS     = 106
	tlangParserDIVEQUALS        = 107
	tlangParserMODEQUALS        = 108
	tlangParserANDEQUALS        = 109
	tlangParserOREQUALS         = 110
	tlangParserIMPLIESEQUALS    = 111
	tlangParserSHREQUALS        = 112
	tlangParserSHLEQUALS        = 113
	tlangParserXOREQUALS        = 114
	tlangParserIDENTIFIER       = 115
	tlangParserWHITESPACE       = 116
	tlangParserBLOCK_COMMENT    = 117
	tlangParserLINE_COMMENT     = 118
)

// tlangParser rules.
const (
	tlangParserRULE_program                = 0
	tlangParserRULE_topLevel               = 1
	tlangParserRULE_procHeader             = 2
	tlangParserRULE_funcHeader             = 3
	tlangParserRULE_subprogramHeader       = 4
	tlangParserRULE_subprogramDeclaration  = 5
	tlangParserRULE_subprogramPrefix       = 6
	tlangParserRULE_classDeclaration       = 7
	tlangParserRULE_classBody              = 8
	tlangParserRULE_exportStatement        = 9
	tlangParserRULE_exportListItem         = 10
	tlangParserRULE_exportList             = 11
	tlangParserRULE_inheritStatement       = 12
	tlangParserRULE_implementStatement     = 13
	tlangParserRULE_idOrFileItem           = 14
	tlangParserRULE_howExport              = 15
	tlangParserRULE_paramDeclaration       = 16
	tlangParserRULE_paramDeclarationList   = 17
	tlangParserRULE_procBody               = 18
	tlangParserRULE_statementOrDeclaration = 19
	tlangParserRULE_primaryExpression      = 20
	tlangParserRULE_expression             = 21
	tlangParserRULE_expressionList         = 22
	tlangParserRULE_declaration            = 23
	tlangParserRULE_statement              = 24
	tlangParserRULE_assignmentOperator     = 25
	tlangParserRULE_putStatement           = 26
	tlangParserRULE_putItem                = 27
	tlangParserRULE_putItemList            = 28
	tlangParserRULE_beginStatement         = 29
	tlangParserRULE_resultStatement        = 30
	tlangParserRULE_newStatement           = 31
	tlangParserRULE_freeStatement          = 32
	tlangParserRULE_forkStatement          = 33
	tlangParserRULE_typeDeclaration        = 34
	tlangParserRULE_basicType              = 35
	tlangParserRULE_referenceType          = 36
	tlangParserRULE_typeSpec               = 37
	tlangParserRULE_indexType              = 38
	tlangParserRULE_indexTypeList          = 39
	tlangParserRULE_stringType             = 40
	tlangParserRULE_recordType             = 41
	tlangParserRULE_recordField            = 42
	tlangParserRULE_variableDeclaration    = 43
	tlangParserRULE_variableIdentifierList = 44
	tlangParserRULE_variableIdentifier     = 45
	tlangParserRULE_arrayDeclaration       = 46
	tlangParserRULE_identifierList         = 47
	tlangParserRULE_literal                = 48
	tlangParserRULE_integer_literal        = 49
	tlangParserRULE_comment                = 50
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
	p.RuleIndex = tlangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_program

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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tlangParserRULE_program)
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserSTRING_LITERAL)|(1<<tlangParserOCTAL_LITERAL)|(1<<tlangParserDECIMAL_LITERAL)|(1<<tlangParserHEX_LITERAL)|(1<<tlangParserBINARY_LITERAL)|(1<<tlangParserREAL_LITERAL)|(1<<tlangParserTYPE)|(1<<tlangParserVAR)|(1<<tlangParserPROCEDURE)|(1<<tlangParserFUNCTION)|(1<<tlangParserCLASS)|(1<<tlangParserPROCESS)|(1<<tlangParserEXIT)|(1<<tlangParserBEGIN)|(1<<tlangParserRETURN)|(1<<tlangParserRESULT)|(1<<tlangParserNEW))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(tlangParserFREE-32))|(1<<(tlangParserFORK-32))|(1<<(tlangParserPUT-32))|(1<<(tlangParserARRAY-32)))) != 0) || (((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tlangParserDEFERRED-68))|(1<<(tlangParserFORWARD-68))|(1<<(tlangParserBODY-68))|(1<<(tlangParserNOT-68))|(1<<(tlangParserCARET-68))|(1<<(tlangParserL_PAREN-68))|(1<<(tlangParserCHEAT-68))|(1<<(tlangParserPLUS-68))|(1<<(tlangParserMINUS-68)))) != 0) || (((_la-115)&-(0x1f+1)) == 0 && ((1<<uint((_la-115)))&((1<<(tlangParserIDENTIFIER-115))|(1<<(tlangParserBLOCK_COMMENT-115))|(1<<(tlangParserLINE_COMMENT-115)))) != 0) {
		{
			p.SetState(102)
			p.TopLevel()
		}

		p.SetState(105)
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
	p.RuleIndex = tlangParserRULE_topLevel
	return p
}

func (*TopLevelContext) IsTopLevelContext() {}

func NewTopLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelContext {
	var p = new(TopLevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_topLevel

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

func (s *TopLevelContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *TopLevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterTopLevel(s)
	}
}

func (s *TopLevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitTopLevel(s)
	}
}

func (s *TopLevelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitTopLevel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) TopLevel() (localctx ITopLevelContext) {
	localctx = NewTopLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tlangParserRULE_topLevel)

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

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.StatementOrDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.ClassDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.Comment()
		}

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
	p.RuleIndex = tlangParserRULE_procHeader
	return p
}

func (*ProcHeaderContext) IsProcHeaderContext() {}

func NewProcHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcHeaderContext {
	var p = new(ProcHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_procHeader

	return p
}

func (s *ProcHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcHeaderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *ProcHeaderContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(tlangParserPROCEDURE, 0)
}

func (s *ProcHeaderContext) PROCESS() antlr.TerminalNode {
	return s.GetToken(tlangParserPROCESS, 0)
}

func (s *ProcHeaderContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *ProcHeaderContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterProcHeader(s)
	}
}

func (s *ProcHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitProcHeader(s)
	}
}

func (s *ProcHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitProcHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ProcHeader() (localctx IProcHeaderContext) {
	localctx = NewProcHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tlangParserRULE_procHeader)
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
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tlangParserPROCEDURE || _la == tlangParserPROCESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(113)
		p.Match(tlangParserIDENTIFIER)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(114)
			p.Match(tlangParserL_PAREN)
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserVAR)|(1<<tlangParserPROCEDURE)|(1<<tlangParserFUNCTION)|(1<<tlangParserPROCESS))) != 0) || (((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tlangParserDEFERRED-68))|(1<<(tlangParserFORWARD-68))|(1<<(tlangParserBODY-68)))) != 0) || _la == tlangParserIDENTIFIER {
			{
				p.SetState(115)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(118)
			p.Match(tlangParserR_PAREN)
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
	p.RuleIndex = tlangParserRULE_funcHeader
	return p
}

func (*FuncHeaderContext) IsFuncHeaderContext() {}

func NewFuncHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncHeaderContext {
	var p = new(FuncHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_funcHeader

	return p
}

func (s *FuncHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncHeaderContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(tlangParserFUNCTION, 0)
}

func (s *FuncHeaderContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(tlangParserIDENTIFIER)
}

func (s *FuncHeaderContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, i)
}

func (s *FuncHeaderContext) COLON() antlr.TerminalNode {
	return s.GetToken(tlangParserCOLON, 0)
}

func (s *FuncHeaderContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FuncHeaderContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *FuncHeaderContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterFuncHeader(s)
	}
}

func (s *FuncHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitFuncHeader(s)
	}
}

func (s *FuncHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitFuncHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) FuncHeader() (localctx IFuncHeaderContext) {
	localctx = NewFuncHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tlangParserRULE_funcHeader)
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
		p.SetState(121)
		p.Match(tlangParserFUNCTION)
	}
	{
		p.SetState(122)
		p.Match(tlangParserIDENTIFIER)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tlangParserL_PAREN {
		{
			p.SetState(123)
			p.Match(tlangParserL_PAREN)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserVAR)|(1<<tlangParserPROCEDURE)|(1<<tlangParserFUNCTION)|(1<<tlangParserPROCESS))) != 0) || (((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tlangParserDEFERRED-68))|(1<<(tlangParserFORWARD-68))|(1<<(tlangParserBODY-68)))) != 0) || _la == tlangParserIDENTIFIER {
			{
				p.SetState(124)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(127)
			p.Match(tlangParserR_PAREN)
		}

	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tlangParserIDENTIFIER {
		{
			p.SetState(130)
			p.Match(tlangParserIDENTIFIER)
		}

	}
	{
		p.SetState(133)
		p.Match(tlangParserCOLON)
	}
	{
		p.SetState(134)
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
	p.RuleIndex = tlangParserRULE_subprogramHeader
	return p
}

func (*SubprogramHeaderContext) IsSubprogramHeaderContext() {}

func NewSubprogramHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubprogramHeaderContext {
	var p = new(SubprogramHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_subprogramHeader

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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterSubprogramHeader(s)
	}
}

func (s *SubprogramHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitSubprogramHeader(s)
	}
}

func (s *SubprogramHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitSubprogramHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) SubprogramHeader() (localctx ISubprogramHeaderContext) {
	localctx = NewSubprogramHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tlangParserRULE_subprogramHeader)
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
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tlangParserDEFERRED-68))|(1<<(tlangParserFORWARD-68))|(1<<(tlangParserBODY-68)))) != 0 {
		{
			p.SetState(136)
			p.SubprogramPrefix()
		}

	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserFUNCTION:
		{
			p.SetState(139)
			p.FuncHeader()
		}

	case tlangParserPROCEDURE, tlangParserPROCESS:
		{
			p.SetState(140)
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
	p.RuleIndex = tlangParserRULE_subprogramDeclaration
	return p
}

func (*SubprogramDeclarationContext) IsSubprogramDeclarationContext() {}

func NewSubprogramDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubprogramDeclarationContext {
	var p = new(SubprogramDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_subprogramDeclaration

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
	return s.GetToken(tlangParserEND, 0)
}

func (s *SubprogramDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *SubprogramDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubprogramDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubprogramDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterSubprogramDeclaration(s)
	}
}

func (s *SubprogramDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitSubprogramDeclaration(s)
	}
}

func (s *SubprogramDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitSubprogramDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) SubprogramDeclaration() (localctx ISubprogramDeclarationContext) {
	localctx = NewSubprogramDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tlangParserRULE_subprogramDeclaration)

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
		p.SetState(143)
		p.SubprogramHeader()
	}
	{
		p.SetState(144)
		p.ProcBody()
	}
	{
		p.SetState(145)
		p.Match(tlangParserEND)
	}
	{
		p.SetState(146)
		p.Match(tlangParserIDENTIFIER)
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
	p.RuleIndex = tlangParserRULE_subprogramPrefix
	return p
}

func (*SubprogramPrefixContext) IsSubprogramPrefixContext() {}

func NewSubprogramPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubprogramPrefixContext {
	var p = new(SubprogramPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_subprogramPrefix

	return p
}

func (s *SubprogramPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SubprogramPrefixContext) DEFERRED() antlr.TerminalNode {
	return s.GetToken(tlangParserDEFERRED, 0)
}

func (s *SubprogramPrefixContext) BODY() antlr.TerminalNode {
	return s.GetToken(tlangParserBODY, 0)
}

func (s *SubprogramPrefixContext) FORWARD() antlr.TerminalNode {
	return s.GetToken(tlangParserFORWARD, 0)
}

func (s *SubprogramPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubprogramPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubprogramPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterSubprogramPrefix(s)
	}
}

func (s *SubprogramPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitSubprogramPrefix(s)
	}
}

func (s *SubprogramPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitSubprogramPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) SubprogramPrefix() (localctx ISubprogramPrefixContext) {
	localctx = NewSubprogramPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tlangParserRULE_subprogramPrefix)
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
		p.SetState(148)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tlangParserDEFERRED-68))|(1<<(tlangParserFORWARD-68))|(1<<(tlangParserBODY-68)))) != 0) {
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
	p.RuleIndex = tlangParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(tlangParserCLASS, 0)
}

func (s *ClassDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(tlangParserIDENTIFIER)
}

func (s *ClassDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, i)
}

func (s *ClassDeclarationContext) END() antlr.TerminalNode {
	return s.GetToken(tlangParserEND, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitClassDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tlangParserRULE_classDeclaration)
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
		p.SetState(150)
		p.Match(tlangParserCLASS)
	}
	{
		p.SetState(151)
		p.Match(tlangParserIDENTIFIER)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserSTRING_LITERAL)|(1<<tlangParserOCTAL_LITERAL)|(1<<tlangParserDECIMAL_LITERAL)|(1<<tlangParserHEX_LITERAL)|(1<<tlangParserBINARY_LITERAL)|(1<<tlangParserREAL_LITERAL)|(1<<tlangParserTYPE)|(1<<tlangParserVAR)|(1<<tlangParserPROCEDURE)|(1<<tlangParserFUNCTION)|(1<<tlangParserCLASS)|(1<<tlangParserPROCESS)|(1<<tlangParserEXIT)|(1<<tlangParserBEGIN)|(1<<tlangParserRETURN)|(1<<tlangParserRESULT)|(1<<tlangParserNEW))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(tlangParserFREE-32))|(1<<(tlangParserFORK-32))|(1<<(tlangParserEXPORT-32))|(1<<(tlangParserINHERIT-32))|(1<<(tlangParserIMPLEMENT-32))|(1<<(tlangParserPUT-32))|(1<<(tlangParserARRAY-32)))) != 0) || (((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tlangParserDEFERRED-68))|(1<<(tlangParserFORWARD-68))|(1<<(tlangParserBODY-68))|(1<<(tlangParserNOT-68))|(1<<(tlangParserCARET-68))|(1<<(tlangParserL_PAREN-68))|(1<<(tlangParserCHEAT-68))|(1<<(tlangParserPLUS-68))|(1<<(tlangParserMINUS-68)))) != 0) || _la == tlangParserIDENTIFIER {
		{
			p.SetState(152)
			p.ClassBody()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(tlangParserEND)
	}
	{
		p.SetState(158)
		p.Match(tlangParserIDENTIFIER)
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
	p.RuleIndex = tlangParserRULE_classBody
	return p
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_classBody

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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (s *ClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitClassBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ClassBody() (localctx IClassBodyContext) {
	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tlangParserRULE_classBody)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserEXPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.ExportStatement()
		}

	case tlangParserINHERIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.InheritStatement()
		}

	case tlangParserIMPLEMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.ImplementStatement()
		}

	case tlangParserSTRING_LITERAL, tlangParserOCTAL_LITERAL, tlangParserDECIMAL_LITERAL, tlangParserHEX_LITERAL, tlangParserBINARY_LITERAL, tlangParserREAL_LITERAL, tlangParserTYPE, tlangParserVAR, tlangParserPROCEDURE, tlangParserFUNCTION, tlangParserCLASS, tlangParserPROCESS, tlangParserEXIT, tlangParserBEGIN, tlangParserRETURN, tlangParserRESULT, tlangParserNEW, tlangParserFREE, tlangParserFORK, tlangParserPUT, tlangParserARRAY, tlangParserDEFERRED, tlangParserFORWARD, tlangParserBODY, tlangParserNOT, tlangParserCARET, tlangParserL_PAREN, tlangParserCHEAT, tlangParserPLUS, tlangParserMINUS, tlangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(163)
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
	p.RuleIndex = tlangParserRULE_exportStatement
	return p
}

func (*ExportStatementContext) IsExportStatementContext() {}

func NewExportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportStatementContext {
	var p = new(ExportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_exportStatement

	return p
}

func (s *ExportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportStatementContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(tlangParserEXPORT, 0)
}

func (s *ExportStatementContext) ExportList() IExportListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportListContext)
}

func (s *ExportStatementContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *ExportStatementContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
}

func (s *ExportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterExportStatement(s)
	}
}

func (s *ExportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitExportStatement(s)
	}
}

func (s *ExportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitExportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ExportStatement() (localctx IExportStatementContext) {
	localctx = NewExportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tlangParserRULE_exportStatement)

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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Match(tlangParserEXPORT)
		}
		{
			p.SetState(167)
			p.exportList(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Match(tlangParserEXPORT)
		}
		{
			p.SetState(169)
			p.Match(tlangParserL_PAREN)
		}
		{
			p.SetState(170)
			p.exportList(0)
		}
		{
			p.SetState(171)
			p.Match(tlangParserR_PAREN)
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
	p.RuleIndex = tlangParserRULE_exportListItem
	return p
}

func (*ExportListItemContext) IsExportListItemContext() {}

func NewExportListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportListItemContext {
	var p = new(ExportListItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_exportListItem

	return p
}

func (s *ExportListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportListItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterExportListItem(s)
	}
}

func (s *ExportListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitExportListItem(s)
	}
}

func (s *ExportListItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitExportListItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ExportListItem() (localctx IExportListItemContext) {
	localctx = NewExportListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tlangParserRULE_exportListItem)
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
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserT__0)|(1<<tlangParserT__1)|(1<<tlangParserT__2)|(1<<tlangParserVAR))) != 0 {
		{
			p.SetState(175)
			p.HowExport()
		}

	}
	{
		p.SetState(178)
		p.Match(tlangParserIDENTIFIER)
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
	p.RuleIndex = tlangParserRULE_exportList
	return p
}

func (*ExportListContext) IsExportListContext() {}

func NewExportListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportListContext {
	var p = new(ExportListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_exportList

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
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *ExportListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterExportList(s)
	}
}

func (s *ExportListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitExportList(s)
	}
}

func (s *ExportListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitExportList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ExportList() (localctx IExportListContext) {
	return p.exportList(0)
}

func (p *tlangParser) exportList(_p int) (localctx IExportListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExportListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExportListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, tlangParserRULE_exportList, _p)

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
		p.SetState(181)
		p.ExportListItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExportListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_exportList)
			p.SetState(183)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(184)
				p.Match(tlangParserCOMMA)
			}
			{
				p.SetState(185)
				p.ExportListItem()
			}

		}
		p.SetState(190)
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
	p.RuleIndex = tlangParserRULE_inheritStatement
	return p
}

func (*InheritStatementContext) IsInheritStatementContext() {}

func NewInheritStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InheritStatementContext {
	var p = new(InheritStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_inheritStatement

	return p
}

func (s *InheritStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InheritStatementContext) INHERIT() antlr.TerminalNode {
	return s.GetToken(tlangParserINHERIT, 0)
}

func (s *InheritStatementContext) IdOrFileItem() IIdOrFileItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdOrFileItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdOrFileItemContext)
}

func (s *InheritStatementContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *InheritStatementContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
}

func (s *InheritStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InheritStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterInheritStatement(s)
	}
}

func (s *InheritStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitInheritStatement(s)
	}
}

func (s *InheritStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitInheritStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) InheritStatement() (localctx IInheritStatementContext) {
	localctx = NewInheritStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tlangParserRULE_inheritStatement)

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

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.Match(tlangParserINHERIT)
		}
		{
			p.SetState(192)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(tlangParserINHERIT)
		}
		{
			p.SetState(194)
			p.Match(tlangParserL_PAREN)
		}
		{
			p.SetState(195)
			p.IdOrFileItem()
		}
		{
			p.SetState(196)
			p.Match(tlangParserR_PAREN)
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
	p.RuleIndex = tlangParserRULE_implementStatement
	return p
}

func (*ImplementStatementContext) IsImplementStatementContext() {}

func NewImplementStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementStatementContext {
	var p = new(ImplementStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_implementStatement

	return p
}

func (s *ImplementStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementStatementContext) IMPLEMENT() antlr.TerminalNode {
	return s.GetToken(tlangParserIMPLEMENT, 0)
}

func (s *ImplementStatementContext) IdOrFileItem() IIdOrFileItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdOrFileItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdOrFileItemContext)
}

func (s *ImplementStatementContext) BY() antlr.TerminalNode {
	return s.GetToken(tlangParserBY, 0)
}

func (s *ImplementStatementContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *ImplementStatementContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
}

func (s *ImplementStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterImplementStatement(s)
	}
}

func (s *ImplementStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitImplementStatement(s)
	}
}

func (s *ImplementStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitImplementStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ImplementStatement() (localctx IImplementStatementContext) {
	localctx = NewImplementStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, tlangParserRULE_implementStatement)
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

	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.Match(tlangParserIMPLEMENT)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tlangParserBY {
			{
				p.SetState(201)
				p.Match(tlangParserBY)
			}

		}
		{
			p.SetState(204)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Match(tlangParserIMPLEMENT)
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tlangParserBY {
			{
				p.SetState(206)
				p.Match(tlangParserBY)
			}

		}
		{
			p.SetState(209)
			p.Match(tlangParserL_PAREN)
		}
		{
			p.SetState(210)
			p.IdOrFileItem()
		}
		{
			p.SetState(211)
			p.Match(tlangParserR_PAREN)
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
	p.RuleIndex = tlangParserRULE_idOrFileItem
	return p
}

func (*IdOrFileItemContext) IsIdOrFileItemContext() {}

func NewIdOrFileItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdOrFileItemContext {
	var p = new(IdOrFileItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_idOrFileItem

	return p
}

func (s *IdOrFileItemContext) GetParser() antlr.Parser { return s.parser }

func (s *IdOrFileItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *IdOrFileItemContext) IN() antlr.TerminalNode {
	return s.GetToken(tlangParserIN, 0)
}

func (s *IdOrFileItemContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(tlangParserSTRING_LITERAL, 0)
}

func (s *IdOrFileItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdOrFileItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdOrFileItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterIdOrFileItem(s)
	}
}

func (s *IdOrFileItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitIdOrFileItem(s)
	}
}

func (s *IdOrFileItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitIdOrFileItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) IdOrFileItem() (localctx IIdOrFileItemContext) {
	localctx = NewIdOrFileItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tlangParserRULE_idOrFileItem)

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

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Match(tlangParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.Match(tlangParserIDENTIFIER)
		}
		{
			p.SetState(217)
			p.Match(tlangParserIN)
		}
		{
			p.SetState(218)
			p.Match(tlangParserSTRING_LITERAL)
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
	p.RuleIndex = tlangParserRULE_howExport
	return p
}

func (*HowExportContext) IsHowExportContext() {}

func NewHowExportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HowExportContext {
	var p = new(HowExportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_howExport

	return p
}

func (s *HowExportContext) GetParser() antlr.Parser { return s.parser }

func (s *HowExportContext) VAR() antlr.TerminalNode {
	return s.GetToken(tlangParserVAR, 0)
}

func (s *HowExportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HowExportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HowExportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterHowExport(s)
	}
}

func (s *HowExportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitHowExport(s)
	}
}

func (s *HowExportContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitHowExport(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) HowExport() (localctx IHowExportContext) {
	localctx = NewHowExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, tlangParserRULE_howExport)
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
		p.SetState(221)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserT__0)|(1<<tlangParserT__1)|(1<<tlangParserT__2)|(1<<tlangParserVAR))) != 0) {
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
	p.RuleIndex = tlangParserRULE_paramDeclaration
	return p
}

func (*ParamDeclarationContext) IsParamDeclarationContext() {}

func NewParamDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclarationContext {
	var p = new(ParamDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_paramDeclaration

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
	return s.GetToken(tlangParserCOLON, 0)
}

func (s *ParamDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ParamDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(tlangParserVAR, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterParamDeclaration(s)
	}
}

func (s *ParamDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitParamDeclaration(s)
	}
}

func (s *ParamDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitParamDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ParamDeclaration() (localctx IParamDeclarationContext) {
	localctx = NewParamDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, tlangParserRULE_paramDeclaration)
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

	p.SetState(231)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserVAR, tlangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tlangParserVAR {
			{
				p.SetState(223)
				p.Match(tlangParserVAR)
			}

		}
		{
			p.SetState(226)
			p.identifierList(0)
		}
		{
			p.SetState(227)
			p.Match(tlangParserCOLON)
		}
		{
			p.SetState(228)
			p.TypeSpec()
		}

	case tlangParserPROCEDURE, tlangParserFUNCTION, tlangParserPROCESS, tlangParserDEFERRED, tlangParserFORWARD, tlangParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
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
	p.RuleIndex = tlangParserRULE_paramDeclarationList
	return p
}

func (*ParamDeclarationListContext) IsParamDeclarationListContext() {}

func NewParamDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclarationListContext {
	var p = new(ParamDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_paramDeclarationList

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
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *ParamDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterParamDeclarationList(s)
	}
}

func (s *ParamDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitParamDeclarationList(s)
	}
}

func (s *ParamDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitParamDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ParamDeclarationList() (localctx IParamDeclarationListContext) {
	return p.paramDeclarationList(0)
}

func (p *tlangParser) paramDeclarationList(_p int) (localctx IParamDeclarationListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParamDeclarationListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParamDeclarationListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, tlangParserRULE_paramDeclarationList, _p)

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
		p.SetState(234)
		p.ParamDeclaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParamDeclarationListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_paramDeclarationList)
			p.SetState(236)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(237)
				p.Match(tlangParserCOMMA)
			}
			{
				p.SetState(238)
				p.ParamDeclaration()
			}

		}
		p.SetState(243)
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
	p.RuleIndex = tlangParserRULE_procBody
	return p
}

func (*ProcBodyContext) IsProcBodyContext() {}

func NewProcBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcBodyContext {
	var p = new(ProcBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_procBody

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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterProcBody(s)
	}
}

func (s *ProcBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitProcBody(s)
	}
}

func (s *ProcBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitProcBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ProcBody() (localctx IProcBodyContext) {
	localctx = NewProcBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tlangParserRULE_procBody)
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
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserSTRING_LITERAL)|(1<<tlangParserOCTAL_LITERAL)|(1<<tlangParserDECIMAL_LITERAL)|(1<<tlangParserHEX_LITERAL)|(1<<tlangParserBINARY_LITERAL)|(1<<tlangParserREAL_LITERAL)|(1<<tlangParserTYPE)|(1<<tlangParserVAR)|(1<<tlangParserPROCEDURE)|(1<<tlangParserFUNCTION)|(1<<tlangParserCLASS)|(1<<tlangParserPROCESS)|(1<<tlangParserEXIT)|(1<<tlangParserBEGIN)|(1<<tlangParserRETURN)|(1<<tlangParserRESULT)|(1<<tlangParserNEW))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(tlangParserFREE-32))|(1<<(tlangParserFORK-32))|(1<<(tlangParserPUT-32))|(1<<(tlangParserARRAY-32)))) != 0) || (((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tlangParserDEFERRED-68))|(1<<(tlangParserFORWARD-68))|(1<<(tlangParserBODY-68))|(1<<(tlangParserNOT-68))|(1<<(tlangParserCARET-68))|(1<<(tlangParserL_PAREN-68))|(1<<(tlangParserCHEAT-68))|(1<<(tlangParserPLUS-68))|(1<<(tlangParserMINUS-68)))) != 0) || _la == tlangParserIDENTIFIER {
		{
			p.SetState(244)
			p.StatementOrDeclaration()
		}

		p.SetState(247)
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
	p.RuleIndex = tlangParserRULE_statementOrDeclaration
	return p
}

func (*StatementOrDeclarationContext) IsStatementOrDeclarationContext() {}

func NewStatementOrDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementOrDeclarationContext {
	var p = new(StatementOrDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_statementOrDeclaration

	return p
}

func (s *StatementOrDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementOrDeclarationContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterStatementOrDeclaration(s)
	}
}

func (s *StatementOrDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitStatementOrDeclaration(s)
	}
}

func (s *StatementOrDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitStatementOrDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) StatementOrDeclaration() (localctx IStatementOrDeclarationContext) {
	localctx = NewStatementOrDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, tlangParserRULE_statementOrDeclaration)

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

	p.SetState(251)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserSTRING_LITERAL, tlangParserOCTAL_LITERAL, tlangParserDECIMAL_LITERAL, tlangParserHEX_LITERAL, tlangParserBINARY_LITERAL, tlangParserREAL_LITERAL, tlangParserEXIT, tlangParserBEGIN, tlangParserRETURN, tlangParserRESULT, tlangParserNEW, tlangParserFREE, tlangParserFORK, tlangParserPUT, tlangParserNOT, tlangParserCARET, tlangParserL_PAREN, tlangParserCHEAT, tlangParserPLUS, tlangParserMINUS, tlangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Statement()
		}

	case tlangParserTYPE, tlangParserVAR, tlangParserPROCEDURE, tlangParserFUNCTION, tlangParserCLASS, tlangParserPROCESS, tlangParserARRAY, tlangParserDEFERRED, tlangParserFORWARD, tlangParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
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
	p.RuleIndex = tlangParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_primaryExpression

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
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *PrimaryExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, tlangParserRULE_primaryExpression)

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

	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserSTRING_LITERAL, tlangParserOCTAL_LITERAL, tlangParserDECIMAL_LITERAL, tlangParserHEX_LITERAL, tlangParserBINARY_LITERAL, tlangParserREAL_LITERAL, tlangParserPLUS, tlangParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Literal()
		}

	case tlangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(tlangParserIDENTIFIER)
		}

	case tlangParserL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.Match(tlangParserL_PAREN)
		}
		{
			p.SetState(256)
			p.expression(0)
		}
		{
			p.SetState(257)
			p.Match(tlangParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token)

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	prefix antlr.Token
	bop    antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetPrefix() antlr.Token { return s.prefix }

func (s *ExpressionContext) GetBop() antlr.Token { return s.bop }

func (s *ExpressionContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *ExpressionContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(tlangParserCARET, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(tlangParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(tlangParserMINUS, 0)
}

func (s *ExpressionContext) CHEAT() antlr.TerminalNode {
	return s.GetToken(tlangParserCHEAT, 0)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(tlangParserNOT, 0)
}

func (s *ExpressionContext) EXP() antlr.TerminalNode {
	return s.GetToken(tlangParserEXP, 0)
}

func (s *ExpressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(tlangParserMULTIPLY, 0)
}

func (s *ExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(tlangParserDIVIDE, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(tlangParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(tlangParserMOD, 0)
}

func (s *ExpressionContext) REM() antlr.TerminalNode {
	return s.GetToken(tlangParserREM, 0)
}

func (s *ExpressionContext) SHR() antlr.TerminalNode {
	return s.GetToken(tlangParserSHR, 0)
}

func (s *ExpressionContext) SHL() antlr.TerminalNode {
	return s.GetToken(tlangParserSHL, 0)
}

func (s *ExpressionContext) XOR() antlr.TerminalNode {
	return s.GetToken(tlangParserXOR, 0)
}

func (s *ExpressionContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(tlangParserLESSTHAN, 0)
}

func (s *ExpressionContext) GREATERTHAN() antlr.TerminalNode {
	return s.GetToken(tlangParserGREATERTHAN, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tlangParserEQUAL, 0)
}

func (s *ExpressionContext) LESSTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(tlangParserLESSTHANEQUAL, 0)
}

func (s *ExpressionContext) GREATERTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(tlangParserGREATERTHANEQUAL, 0)
}

func (s *ExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(tlangParserNOTEQUAL, 0)
}

func (s *ExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(tlangParserIN, 0)
}

func (s *ExpressionContext) NOT_IN() antlr.TerminalNode {
	return s.GetToken(tlangParserNOT_IN, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(tlangParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(tlangParserOR, 0)
}

func (s *ExpressionContext) IMPLIES() antlr.TerminalNode {
	return s.GetToken(tlangParserIMPLIES, 0)
}

func (s *ExpressionContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(tlangParserASSIGNMENT, 0)
}

func (s *ExpressionContext) PLUSEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserPLUSEQUALS, 0)
}

func (s *ExpressionContext) MINUSEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserMINUSEQUALS, 0)
}

func (s *ExpressionContext) MULTIPLYEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserMULTIPLYEQUALS, 0)
}

func (s *ExpressionContext) DIVIDEEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserDIVIDEEQUALS, 0)
}

func (s *ExpressionContext) DIVEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserDIVEQUALS, 0)
}

func (s *ExpressionContext) SHLEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserSHLEQUALS, 0)
}

func (s *ExpressionContext) SHREQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserSHREQUALS, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *ExpressionContext) DEREFERENCE() antlr.TerminalNode {
	return s.GetToken(tlangParserDEREFERENCE, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(tlangParserDOT, 0)
}

func (s *ExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *ExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
}

func (s *ExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *tlangParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, tlangParserRULE_expression, _p)
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
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(262)
			p.PrimaryExpression()
		}

	case 2:
		{
			p.SetState(263)

			var _m = p.Match(tlangParserCARET)

			localctx.(*ExpressionContext).prefix = _m
		}
		{
			p.SetState(264)
			p.expression(13)
		}

	case 3:
		{
			p.SetState(265)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-79)&-(0x1f+1)) == 0 && ((1<<uint((_la-79)))&((1<<(tlangParserCHEAT-79))|(1<<(tlangParserPLUS-79))|(1<<(tlangParserMINUS-79)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(266)
			p.expression(9)
		}

	case 4:
		{
			p.SetState(267)

			var _m = p.Match(tlangParserNOT)

			localctx.(*ExpressionContext).prefix = _m
		}
		{
			p.SetState(268)
			p.expression(5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(272)

					var _m = p.Match(tlangParserEXP)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(273)
					p.expression(12)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(275)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-84)&-(0x1f+1)) == 0 && ((1<<uint((_la-84)))&((1<<(tlangParserMULTIPLY-84))|(1<<(tlangParserDIVIDE-84))|(1<<(tlangParserDIV-84))|(1<<(tlangParserMOD-84))|(1<<(tlangParserREM-84))|(1<<(tlangParserSHR-84))|(1<<(tlangParserSHL-84)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(276)
					p.expression(9)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(278)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-82)&-(0x1f+1)) == 0 && ((1<<uint((_la-82)))&((1<<(tlangParserPLUS-82))|(1<<(tlangParserMINUS-82))|(1<<(tlangParserXOR-82)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(279)
					p.expression(8)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(281)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == tlangParserNOT_IN || (((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(tlangParserLESSTHAN-90))|(1<<(tlangParserGREATERTHAN-90))|(1<<(tlangParserEQUAL-90))|(1<<(tlangParserLESSTHANEQUAL-90))|(1<<(tlangParserGREATERTHANEQUAL-90))|(1<<(tlangParserNOTEQUAL-90))|(1<<(tlangParserIN-90)))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(282)
					p.expression(7)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(284)

					var _m = p.Match(tlangParserAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(285)
					p.expression(5)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(287)

					var _m = p.Match(tlangParserOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(288)
					p.expression(4)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(290)

					var _m = p.Match(tlangParserIMPLIES)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(291)
					p.expression(3)
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(293)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la-81)&-(0x1f+1)) == 0 && ((1<<uint((_la-81)))&((1<<(tlangParserASSIGNMENT-81))|(1<<(tlangParserPLUSEQUALS-81))|(1<<(tlangParserMINUSEQUALS-81))|(1<<(tlangParserMULTIPLYEQUALS-81))|(1<<(tlangParserDIVIDEEQUALS-81))|(1<<(tlangParserDIVEQUALS-81))|(1<<(tlangParserSHREQUALS-81)))) != 0) || _la == tlangParserSHLEQUALS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(294)
					p.expression(2)
				}

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(296)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == tlangParserDOT || _la == tlangParserDEREFERENCE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(297)
					p.Match(tlangParserIDENTIFIER)
				}

			case 10:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_expression)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(299)
					p.Match(tlangParserL_PAREN)
				}
				p.SetState(301)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tlangParserSTRING_LITERAL)|(1<<tlangParserOCTAL_LITERAL)|(1<<tlangParserDECIMAL_LITERAL)|(1<<tlangParserHEX_LITERAL)|(1<<tlangParserBINARY_LITERAL)|(1<<tlangParserREAL_LITERAL))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(tlangParserNOT-71))|(1<<(tlangParserCARET-71))|(1<<(tlangParserL_PAREN-71))|(1<<(tlangParserCHEAT-71))|(1<<(tlangParserPLUS-71))|(1<<(tlangParserMINUS-71)))) != 0) || _la == tlangParserIDENTIFIER {
					{
						p.SetState(300)
						p.ExpressionList()
					}

				}
				{
					p.SetState(303)
					p.Match(tlangParserR_PAREN)
				}

			}

		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.RuleIndex = tlangParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tlangParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tlangParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, tlangParserRULE_expressionList)
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
		p.SetState(309)
		p.expression(0)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tlangParserCOMMA {
		{
			p.SetState(310)
			p.Match(tlangParserCOMMA)
		}
		{
			p.SetState(311)
			p.expression(0)
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = tlangParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_declaration

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

func (s *DeclarationContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, tlangParserRULE_declaration)

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

	p.SetState(322)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.TypeDeclaration()
		}

	case tlangParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)
			p.VariableDeclaration()
		}

	case tlangParserARRAY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.ArrayDeclaration()
		}

	case tlangParserPROCEDURE, tlangParserFUNCTION, tlangParserPROCESS, tlangParserDEFERRED, tlangParserFORWARD, tlangParserBODY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(320)
			p.SubprogramDeclaration()
		}

	case tlangParserCLASS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(321)
			p.ClassDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) PutStatement() IPutStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPutStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPutStatementContext)
}

func (s *StatementContext) EXIT() antlr.TerminalNode {
	return s.GetToken(tlangParserEXIT, 0)
}

func (s *StatementContext) BeginStatement() IBeginStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeginStatementContext)
}

func (s *StatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(tlangParserRETURN, 0)
}

func (s *StatementContext) ResultStatement() IResultStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResultStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResultStatementContext)
}

func (s *StatementContext) NewStatement() INewStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewStatementContext)
}

func (s *StatementContext) FreeStatement() IFreeStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFreeStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFreeStatementContext)
}

func (s *StatementContext) ForkStatement() IForkStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForkStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForkStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, tlangParserRULE_statement)

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

	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserSTRING_LITERAL, tlangParserOCTAL_LITERAL, tlangParserDECIMAL_LITERAL, tlangParserHEX_LITERAL, tlangParserBINARY_LITERAL, tlangParserREAL_LITERAL, tlangParserNOT, tlangParserCARET, tlangParserL_PAREN, tlangParserCHEAT, tlangParserPLUS, tlangParserMINUS, tlangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.expression(0)
		}

	case tlangParserPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.PutStatement()
		}

	case tlangParserEXIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(326)
			p.Match(tlangParserEXIT)
		}

	case tlangParserBEGIN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(327)
			p.BeginStatement()
		}

	case tlangParserRETURN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(328)
			p.Match(tlangParserRETURN)
		}

	case tlangParserRESULT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(329)
			p.ResultStatement()
		}

	case tlangParserNEW:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(330)
			p.NewStatement()
		}

	case tlangParserFREE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(331)
			p.FreeStatement()
		}

	case tlangParserFORK:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(332)
			p.ForkStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlangParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(tlangParserASSIGNMENT, 0)
}

func (s *AssignmentOperatorContext) PLUSEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserPLUSEQUALS, 0)
}

func (s *AssignmentOperatorContext) MINUSEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserMINUSEQUALS, 0)
}

func (s *AssignmentOperatorContext) MULTIPLYEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserMULTIPLYEQUALS, 0)
}

func (s *AssignmentOperatorContext) DIVIDEEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserDIVIDEEQUALS, 0)
}

func (s *AssignmentOperatorContext) DIVEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserDIVEQUALS, 0)
}

func (s *AssignmentOperatorContext) SHLEQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserSHLEQUALS, 0)
}

func (s *AssignmentOperatorContext) SHREQUALS() antlr.TerminalNode {
	return s.GetToken(tlangParserSHREQUALS, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitAssignmentOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, tlangParserRULE_assignmentOperator)
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
		p.SetState(335)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-81)&-(0x1f+1)) == 0 && ((1<<uint((_la-81)))&((1<<(tlangParserASSIGNMENT-81))|(1<<(tlangParserPLUSEQUALS-81))|(1<<(tlangParserMINUSEQUALS-81))|(1<<(tlangParserMULTIPLYEQUALS-81))|(1<<(tlangParserDIVIDEEQUALS-81))|(1<<(tlangParserDIVEQUALS-81))|(1<<(tlangParserSHREQUALS-81)))) != 0) || _la == tlangParserSHLEQUALS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = tlangParserRULE_putStatement
	return p
}

func (*PutStatementContext) IsPutStatementContext() {}

func NewPutStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutStatementContext {
	var p = new(PutStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_putStatement

	return p
}

func (s *PutStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PutStatementContext) PUT() antlr.TerminalNode {
	return s.GetToken(tlangParserPUT, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterPutStatement(s)
	}
}

func (s *PutStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitPutStatement(s)
	}
}

func (s *PutStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitPutStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) PutStatement() (localctx IPutStatementContext) {
	localctx = NewPutStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, tlangParserRULE_putStatement)

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
		p.SetState(337)
		p.Match(tlangParserPUT)
	}
	{
		p.SetState(338)
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
	p.RuleIndex = tlangParserRULE_putItem
	return p
}

func (*PutItemContext) IsPutItemContext() {}

func NewPutItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutItemContext {
	var p = new(PutItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_putItem

	return p
}

func (s *PutItemContext) GetParser() antlr.Parser { return s.parser }

func (s *PutItemContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *PutItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PutItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterPutItem(s)
	}
}

func (s *PutItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitPutItem(s)
	}
}

func (s *PutItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitPutItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) PutItem() (localctx IPutItemContext) {
	localctx = NewPutItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, tlangParserRULE_putItem)

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
		p.SetState(340)
		p.Statement()
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
	p.RuleIndex = tlangParserRULE_putItemList
	return p
}

func (*PutItemListContext) IsPutItemListContext() {}

func NewPutItemListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutItemListContext {
	var p = new(PutItemListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_putItemList

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
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *PutItemListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutItemListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PutItemListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterPutItemList(s)
	}
}

func (s *PutItemListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitPutItemList(s)
	}
}

func (s *PutItemListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitPutItemList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) PutItemList() (localctx IPutItemListContext) {
	return p.putItemList(0)
}

func (p *tlangParser) putItemList(_p int) (localctx IPutItemListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPutItemListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPutItemListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, tlangParserRULE_putItemList, _p)

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
		p.SetState(343)
		p.PutItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPutItemListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_putItemList)
			p.SetState(345)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(346)
				p.Match(tlangParserCOMMA)
			}
			{
				p.SetState(347)
				p.PutItem()
			}

		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	p.RuleIndex = tlangParserRULE_beginStatement
	return p
}

func (*BeginStatementContext) IsBeginStatementContext() {}

func NewBeginStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginStatementContext {
	var p = new(BeginStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_beginStatement

	return p
}

func (s *BeginStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginStatementContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(tlangParserBEGIN, 0)
}

func (s *BeginStatementContext) StatementOrDeclaration() IStatementOrDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementOrDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementOrDeclarationContext)
}

func (s *BeginStatementContext) END() antlr.TerminalNode {
	return s.GetToken(tlangParserEND, 0)
}

func (s *BeginStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterBeginStatement(s)
	}
}

func (s *BeginStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitBeginStatement(s)
	}
}

func (s *BeginStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitBeginStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) BeginStatement() (localctx IBeginStatementContext) {
	localctx = NewBeginStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, tlangParserRULE_beginStatement)

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
		p.SetState(353)
		p.Match(tlangParserBEGIN)
	}
	{
		p.SetState(354)
		p.StatementOrDeclaration()
	}
	{
		p.SetState(355)
		p.Match(tlangParserEND)
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
	p.RuleIndex = tlangParserRULE_resultStatement
	return p
}

func (*ResultStatementContext) IsResultStatementContext() {}

func NewResultStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultStatementContext {
	var p = new(ResultStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_resultStatement

	return p
}

func (s *ResultStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultStatementContext) RESULT() antlr.TerminalNode {
	return s.GetToken(tlangParserRESULT, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterResultStatement(s)
	}
}

func (s *ResultStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitResultStatement(s)
	}
}

func (s *ResultStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitResultStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ResultStatement() (localctx IResultStatementContext) {
	localctx = NewResultStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, tlangParserRULE_resultStatement)

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
		p.SetState(357)
		p.Match(tlangParserRESULT)
	}
	{
		p.SetState(358)
		p.expression(0)
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
	p.RuleIndex = tlangParserRULE_newStatement
	return p
}

func (*NewStatementContext) IsNewStatementContext() {}

func NewNewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewStatementContext {
	var p = new(NewStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_newStatement

	return p
}

func (s *NewStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *NewStatementContext) NEW() antlr.TerminalNode {
	return s.GetToken(tlangParserNEW, 0)
}

func (s *NewStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(tlangParserIDENTIFIER)
}

func (s *NewStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, i)
}

func (s *NewStatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *NewStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterNewStatement(s)
	}
}

func (s *NewStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitNewStatement(s)
	}
}

func (s *NewStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitNewStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) NewStatement() (localctx INewStatementContext) {
	localctx = NewNewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, tlangParserRULE_newStatement)

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
		p.SetState(360)
		p.Match(tlangParserNEW)
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(361)
			p.Match(tlangParserIDENTIFIER)
		}
		{
			p.SetState(362)
			p.Match(tlangParserCOMMA)
		}

	}
	{
		p.SetState(365)
		p.Match(tlangParserIDENTIFIER)
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
	p.RuleIndex = tlangParserRULE_freeStatement
	return p
}

func (*FreeStatementContext) IsFreeStatementContext() {}

func NewFreeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FreeStatementContext {
	var p = new(FreeStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_freeStatement

	return p
}

func (s *FreeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FreeStatementContext) FREE() antlr.TerminalNode {
	return s.GetToken(tlangParserFREE, 0)
}

func (s *FreeStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(tlangParserIDENTIFIER)
}

func (s *FreeStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, i)
}

func (s *FreeStatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *FreeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FreeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FreeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterFreeStatement(s)
	}
}

func (s *FreeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitFreeStatement(s)
	}
}

func (s *FreeStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitFreeStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) FreeStatement() (localctx IFreeStatementContext) {
	localctx = NewFreeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, tlangParserRULE_freeStatement)

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
		p.SetState(367)
		p.Match(tlangParserFREE)
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(368)
			p.Match(tlangParserIDENTIFIER)
		}
		{
			p.SetState(369)
			p.Match(tlangParserCOMMA)
		}

	}
	{
		p.SetState(372)
		p.Match(tlangParserIDENTIFIER)
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
	p.RuleIndex = tlangParserRULE_forkStatement
	return p
}

func (*ForkStatementContext) IsForkStatementContext() {}

func NewForkStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForkStatementContext {
	var p = new(ForkStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_forkStatement

	return p
}

func (s *ForkStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForkStatementContext) FORK() antlr.TerminalNode {
	return s.GetToken(tlangParserFORK, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterForkStatement(s)
	}
}

func (s *ForkStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitForkStatement(s)
	}
}

func (s *ForkStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitForkStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ForkStatement() (localctx IForkStatementContext) {
	localctx = NewForkStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, tlangParserRULE_forkStatement)

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
		p.SetState(374)
		p.Match(tlangParserFORK)
	}
	{
		p.SetState(375)
		p.expression(0)
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
	p.RuleIndex = tlangParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(tlangParserTYPE, 0)
}

func (s *TypeDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *TypeDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(tlangParserCOLON, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, tlangParserRULE_typeDeclaration)

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
		p.SetState(377)
		p.Match(tlangParserTYPE)
	}
	{
		p.SetState(378)
		p.Match(tlangParserIDENTIFIER)
	}
	{
		p.SetState(379)
		p.Match(tlangParserCOLON)
	}
	{
		p.SetState(380)
		p.TypeSpec()
	}

	return localctx
}

// IBasicTypeContext is an interface to support dynamic dispatch.
type IBasicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicTypeContext differentiates from other interfaces.
	IsBasicTypeContext()
}

type BasicTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicTypeContext() *BasicTypeContext {
	var p = new(BasicTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlangParserRULE_basicType
	return p
}

func (*BasicTypeContext) IsBasicTypeContext() {}

func NewBasicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicTypeContext {
	var p = new(BasicTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_basicType

	return p
}

func (s *BasicTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(tlangParserINT, 0)
}

func (s *BasicTypeContext) REAL() antlr.TerminalNode {
	return s.GetToken(tlangParserREAL, 0)
}

func (s *BasicTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(tlangParserBOOLEAN, 0)
}

func (s *BasicTypeContext) NAT() antlr.TerminalNode {
	return s.GetToken(tlangParserNAT, 0)
}

func (s *BasicTypeContext) INT1() antlr.TerminalNode {
	return s.GetToken(tlangParserINT1, 0)
}

func (s *BasicTypeContext) INT2() antlr.TerminalNode {
	return s.GetToken(tlangParserINT2, 0)
}

func (s *BasicTypeContext) INT4() antlr.TerminalNode {
	return s.GetToken(tlangParserINT4, 0)
}

func (s *BasicTypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(tlangParserINT8, 0)
}

func (s *BasicTypeContext) NAT1() antlr.TerminalNode {
	return s.GetToken(tlangParserNAT1, 0)
}

func (s *BasicTypeContext) NAT2() antlr.TerminalNode {
	return s.GetToken(tlangParserNAT2, 0)
}

func (s *BasicTypeContext) NAT4() antlr.TerminalNode {
	return s.GetToken(tlangParserNAT4, 0)
}

func (s *BasicTypeContext) NAT8() antlr.TerminalNode {
	return s.GetToken(tlangParserNAT8, 0)
}

func (s *BasicTypeContext) REAL4() antlr.TerminalNode {
	return s.GetToken(tlangParserREAL4, 0)
}

func (s *BasicTypeContext) REAL8() antlr.TerminalNode {
	return s.GetToken(tlangParserREAL8, 0)
}

func (s *BasicTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(tlangParserCHAR, 0)
}

func (s *BasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterBasicType(s)
	}
}

func (s *BasicTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitBasicType(s)
	}
}

func (s *BasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) BasicType() (localctx IBasicTypeContext) {
	localctx = NewBasicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, tlangParserRULE_basicType)
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
		p.SetState(382)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(tlangParserINT-47))|(1<<(tlangParserREAL-47))|(1<<(tlangParserBOOLEAN-47))|(1<<(tlangParserNAT-47))|(1<<(tlangParserINT1-47))|(1<<(tlangParserINT2-47))|(1<<(tlangParserINT4-47))|(1<<(tlangParserINT8-47))|(1<<(tlangParserNAT1-47))|(1<<(tlangParserNAT2-47))|(1<<(tlangParserNAT4-47))|(1<<(tlangParserNAT8-47))|(1<<(tlangParserREAL4-47))|(1<<(tlangParserREAL8-47))|(1<<(tlangParserCHAR-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReferenceTypeContext is an interface to support dynamic dispatch.
type IReferenceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceTypeContext differentiates from other interfaces.
	IsReferenceTypeContext()
}

type ReferenceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceTypeContext() *ReferenceTypeContext {
	var p = new(ReferenceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlangParserRULE_referenceType
	return p
}

func (*ReferenceTypeContext) IsReferenceTypeContext() {}

func NewReferenceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceTypeContext {
	var p = new(ReferenceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_referenceType

	return p
}

func (s *ReferenceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *ReferenceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterReferenceType(s)
	}
}

func (s *ReferenceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitReferenceType(s)
	}
}

func (s *ReferenceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitReferenceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ReferenceType() (localctx IReferenceTypeContext) {
	localctx = NewReferenceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, tlangParserRULE_referenceType)

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
		p.SetState(384)
		p.Match(tlangParserIDENTIFIER)
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
	p.RuleIndex = tlangParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) BasicType() IBasicTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
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

func (s *TypeSpecContext) ArrayDeclaration() IArrayDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayDeclarationContext)
}

func (s *TypeSpecContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *TypeSpecContext) ReferenceType() IReferenceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceTypeContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, tlangParserRULE_typeSpec)

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

	p.SetState(393)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserINT, tlangParserREAL, tlangParserBOOLEAN, tlangParserNAT, tlangParserINT1, tlangParserINT2, tlangParserINT4, tlangParserINT8, tlangParserNAT1, tlangParserNAT2, tlangParserNAT4, tlangParserNAT8, tlangParserREAL4, tlangParserREAL8, tlangParserCHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)
			p.BasicType()
		}

	case tlangParserOCTAL_LITERAL, tlangParserDECIMAL_LITERAL, tlangParserHEX_LITERAL, tlangParserBINARY_LITERAL, tlangParserPLUS, tlangParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)
			p.IndexType()
		}

	case tlangParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(388)
			p.StringType()
		}

	case tlangParserRECORD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(389)
			p.RecordType()
		}

	case tlangParserARRAY:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(390)
			p.ArrayDeclaration()
		}

	case tlangParserCLASS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(391)
			p.ClassDeclaration()
		}

	case tlangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(392)
			p.ReferenceType()
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
	p.RuleIndex = tlangParserRULE_indexType
	return p
}

func (*IndexTypeContext) IsIndexTypeContext() {}

func NewIndexTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTypeContext {
	var p = new(IndexTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_indexType

	return p
}

func (s *IndexTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexTypeContext) AllInteger_literal() []IInteger_literalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInteger_literalContext)(nil)).Elem())
	var tst = make([]IInteger_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInteger_literalContext)
		}
	}

	return tst
}

func (s *IndexTypeContext) Integer_literal(i int) IInteger_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_literalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *IndexTypeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(tlangParserRANGE, 0)
}

func (s *IndexTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterIndexType(s)
	}
}

func (s *IndexTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitIndexType(s)
	}
}

func (s *IndexTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitIndexType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) IndexType() (localctx IIndexTypeContext) {
	localctx = NewIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, tlangParserRULE_indexType)

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
		p.SetState(395)
		p.Integer_literal()
	}
	{
		p.SetState(396)
		p.Match(tlangParserRANGE)
	}
	{
		p.SetState(397)
		p.Integer_literal()
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
	p.RuleIndex = tlangParserRULE_indexTypeList
	return p
}

func (*IndexTypeListContext) IsIndexTypeListContext() {}

func NewIndexTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTypeListContext {
	var p = new(IndexTypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_indexTypeList

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
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *IndexTypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterIndexTypeList(s)
	}
}

func (s *IndexTypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitIndexTypeList(s)
	}
}

func (s *IndexTypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitIndexTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) IndexTypeList() (localctx IIndexTypeListContext) {
	return p.indexTypeList(0)
}

func (p *tlangParser) indexTypeList(_p int) (localctx IIndexTypeListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIndexTypeListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIndexTypeListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, tlangParserRULE_indexTypeList, _p)

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
		p.SetState(400)
		p.IndexType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIndexTypeListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_indexTypeList)
			p.SetState(402)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(403)
				p.Match(tlangParserCOMMA)
			}
			{
				p.SetState(404)
				p.IndexType()
			}

		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
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
	p.RuleIndex = tlangParserRULE_stringType
	return p
}

func (*StringTypeContext) IsStringTypeContext() {}

func NewStringTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTypeContext {
	var p = new(StringTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_stringType

	return p
}

func (s *StringTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(tlangParserSTRING, 0)
}

func (s *StringTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserL_PAREN, 0)
}

func (s *StringTypeContext) Integer_literal() IInteger_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *StringTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tlangParserR_PAREN, 0)
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) StringType() (localctx IStringTypeContext) {
	localctx = NewStringTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, tlangParserRULE_stringType)

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
		p.SetState(410)
		p.Match(tlangParserSTRING)
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(411)
			p.Match(tlangParserL_PAREN)
		}
		{
			p.SetState(412)
			p.Integer_literal()
		}
		{
			p.SetState(413)
			p.Match(tlangParserR_PAREN)
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
	p.RuleIndex = tlangParserRULE_recordType
	return p
}

func (*RecordTypeContext) IsRecordTypeContext() {}

func NewRecordTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordTypeContext {
	var p = new(RecordTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_recordType

	return p
}

func (s *RecordTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordTypeContext) AllRECORD() []antlr.TerminalNode {
	return s.GetTokens(tlangParserRECORD)
}

func (s *RecordTypeContext) RECORD(i int) antlr.TerminalNode {
	return s.GetToken(tlangParserRECORD, i)
}

func (s *RecordTypeContext) END() antlr.TerminalNode {
	return s.GetToken(tlangParserEND, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterRecordType(s)
	}
}

func (s *RecordTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitRecordType(s)
	}
}

func (s *RecordTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitRecordType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) RecordType() (localctx IRecordTypeContext) {
	localctx = NewRecordTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, tlangParserRULE_recordType)
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
		p.SetState(417)
		p.Match(tlangParserRECORD)
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tlangParserIDENTIFIER {
		{
			p.SetState(418)
			p.RecordField()
		}

		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(423)
		p.Match(tlangParserEND)
	}
	{
		p.SetState(424)
		p.Match(tlangParserRECORD)
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
	p.RuleIndex = tlangParserRULE_recordField
	return p
}

func (*RecordFieldContext) IsRecordFieldContext() {}

func NewRecordFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordFieldContext {
	var p = new(RecordFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_recordField

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
	return s.GetToken(tlangParserCOLON, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterRecordField(s)
	}
}

func (s *RecordFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitRecordField(s)
	}
}

func (s *RecordFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitRecordField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) RecordField() (localctx IRecordFieldContext) {
	localctx = NewRecordFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, tlangParserRULE_recordField)

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
		p.SetState(426)
		p.identifierList(0)
	}
	{
		p.SetState(427)
		p.Match(tlangParserCOLON)
	}
	{
		p.SetState(428)
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
	p.RuleIndex = tlangParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(tlangParserVAR, 0)
}

func (s *VariableDeclarationContext) VariableIdentifierList() IVariableIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableIdentifierListContext)
}

func (s *VariableDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(tlangParserCOLON, 0)
}

func (s *VariableDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *VariableDeclarationContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(tlangParserASSIGNMENT, 0)
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
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, tlangParserRULE_variableDeclaration)
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

	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(430)
			p.Match(tlangParserVAR)
		}
		{
			p.SetState(431)
			p.VariableIdentifierList()
		}
		{
			p.SetState(432)
			p.Match(tlangParserCOLON)
		}
		{
			p.SetState(433)
			p.TypeSpec()
		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tlangParserASSIGNMENT {
			{
				p.SetState(434)
				p.Match(tlangParserASSIGNMENT)
			}
			{
				p.SetState(435)
				p.expression(0)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Match(tlangParserVAR)
		}
		{
			p.SetState(439)
			p.VariableIdentifierList()
		}
		{
			p.SetState(440)
			p.Match(tlangParserASSIGNMENT)
		}
		{
			p.SetState(441)
			p.expression(0)
		}

	}

	return localctx
}

// IVariableIdentifierListContext is an interface to support dynamic dispatch.
type IVariableIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableIdentifierListContext differentiates from other interfaces.
	IsVariableIdentifierListContext()
}

type VariableIdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableIdentifierListContext() *VariableIdentifierListContext {
	var p = new(VariableIdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlangParserRULE_variableIdentifierList
	return p
}

func (*VariableIdentifierListContext) IsVariableIdentifierListContext() {}

func NewVariableIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableIdentifierListContext {
	var p = new(VariableIdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_variableIdentifierList

	return p
}

func (s *VariableIdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableIdentifierListContext) AllVariableIdentifier() []IVariableIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableIdentifierContext)(nil)).Elem())
	var tst = make([]IVariableIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableIdentifierContext)
		}
	}

	return tst
}

func (s *VariableIdentifierListContext) VariableIdentifier(i int) IVariableIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableIdentifierContext)
}

func (s *VariableIdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tlangParserCOMMA)
}

func (s *VariableIdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tlangParserCOMMA, i)
}

func (s *VariableIdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableIdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableIdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterVariableIdentifierList(s)
	}
}

func (s *VariableIdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitVariableIdentifierList(s)
	}
}

func (s *VariableIdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitVariableIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) VariableIdentifierList() (localctx IVariableIdentifierListContext) {
	localctx = NewVariableIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, tlangParserRULE_variableIdentifierList)
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
		p.SetState(445)
		p.VariableIdentifier()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tlangParserCOMMA {
		{
			p.SetState(446)
			p.Match(tlangParserCOMMA)
		}
		{
			p.SetState(447)
			p.VariableIdentifier()
		}

		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = tlangParserRULE_variableIdentifier
	return p
}

func (*VariableIdentifierContext) IsVariableIdentifierContext() {}

func NewVariableIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableIdentifierContext {
	var p = new(VariableIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_variableIdentifier

	return p
}

func (s *VariableIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *VariableIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterVariableIdentifier(s)
	}
}

func (s *VariableIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitVariableIdentifier(s)
	}
}

func (s *VariableIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitVariableIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) VariableIdentifier() (localctx IVariableIdentifierContext) {
	localctx = NewVariableIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, tlangParserRULE_variableIdentifier)

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
		p.SetState(453)
		p.Match(tlangParserIDENTIFIER)
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
	p.RuleIndex = tlangParserRULE_arrayDeclaration
	return p
}

func (*ArrayDeclarationContext) IsArrayDeclarationContext() {}

func NewArrayDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclarationContext {
	var p = new(ArrayDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_arrayDeclaration

	return p
}

func (s *ArrayDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclarationContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(tlangParserARRAY, 0)
}

func (s *ArrayDeclarationContext) IndexType() IIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeContext)
}

func (s *ArrayDeclarationContext) OF() antlr.TerminalNode {
	return s.GetToken(tlangParserOF, 0)
}

func (s *ArrayDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ArrayDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *ArrayDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterArrayDeclaration(s)
	}
}

func (s *ArrayDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitArrayDeclaration(s)
	}
}

func (s *ArrayDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitArrayDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) ArrayDeclaration() (localctx IArrayDeclarationContext) {
	localctx = NewArrayDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, tlangParserRULE_arrayDeclaration)

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

	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(455)
			p.Match(tlangParserARRAY)
		}
		{
			p.SetState(456)
			p.IndexType()
		}
		{
			p.SetState(457)
			p.Match(tlangParserOF)
		}
		{
			p.SetState(458)
			p.TypeSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)
			p.Match(tlangParserARRAY)
		}
		{
			p.SetState(461)
			p.IndexType()
		}
		{
			p.SetState(462)
			p.Match(tlangParserCOMMA)
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
	p.RuleIndex = tlangParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tlangParserIDENTIFIER, 0)
}

func (s *IdentifierListContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *IdentifierListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(tlangParserCOMMA, 0)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) IdentifierList() (localctx IIdentifierListContext) {
	return p.identifierList(0)
}

func (p *tlangParser) identifierList(_p int) (localctx IIdentifierListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIdentifierListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 94
	p.EnterRecursionRule(localctx, 94, tlangParserRULE_identifierList, _p)

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
		p.SetState(467)
		p.Match(tlangParserIDENTIFIER)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIdentifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, tlangParserRULE_identifierList)
			p.SetState(469)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(470)
				p.Match(tlangParserCOMMA)
			}
			{
				p.SetState(471)
				p.Match(tlangParserIDENTIFIER)
			}

		}
		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
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
	p.RuleIndex = tlangParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(tlangParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(tlangParserREAL_LITERAL, 0)
}

func (s *LiteralContext) Integer_literal() IInteger_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, tlangParserRULE_literal)

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

	p.SetState(480)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)
			p.Match(tlangParserSTRING_LITERAL)
		}

	case tlangParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.Match(tlangParserREAL_LITERAL)
		}

	case tlangParserOCTAL_LITERAL, tlangParserDECIMAL_LITERAL, tlangParserHEX_LITERAL, tlangParserBINARY_LITERAL, tlangParserPLUS, tlangParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(479)
			p.Integer_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInteger_literalContext is an interface to support dynamic dispatch.
type IInteger_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSign returns the sign token.
	GetSign() antlr.Token

	// SetSign sets the sign token.
	SetSign(antlr.Token)

	// IsInteger_literalContext differentiates from other interfaces.
	IsInteger_literalContext()
}

type Integer_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sign   antlr.Token
}

func NewEmptyInteger_literalContext() *Integer_literalContext {
	var p = new(Integer_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlangParserRULE_integer_literal
	return p
}

func (*Integer_literalContext) IsInteger_literalContext() {}

func NewInteger_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_literalContext {
	var p = new(Integer_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_integer_literal

	return p
}

func (s *Integer_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_literalContext) GetSign() antlr.Token { return s.sign }

func (s *Integer_literalContext) SetSign(v antlr.Token) { s.sign = v }

func (s *Integer_literalContext) OCTAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(tlangParserOCTAL_LITERAL, 0)
}

func (s *Integer_literalContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(tlangParserDECIMAL_LITERAL, 0)
}

func (s *Integer_literalContext) PLUS() antlr.TerminalNode {
	return s.GetToken(tlangParserPLUS, 0)
}

func (s *Integer_literalContext) MINUS() antlr.TerminalNode {
	return s.GetToken(tlangParserMINUS, 0)
}

func (s *Integer_literalContext) BINARY_LITERAL() antlr.TerminalNode {
	return s.GetToken(tlangParserBINARY_LITERAL, 0)
}

func (s *Integer_literalContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(tlangParserHEX_LITERAL, 0)
}

func (s *Integer_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterInteger_literal(s)
	}
}

func (s *Integer_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitInteger_literal(s)
	}
}

func (s *Integer_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitInteger_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) Integer_literal() (localctx IInteger_literalContext) {
	localctx = NewInteger_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, tlangParserRULE_integer_literal)
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

	p.SetState(489)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tlangParserOCTAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(482)
			p.Match(tlangParserOCTAL_LITERAL)
		}

	case tlangParserDECIMAL_LITERAL, tlangParserPLUS, tlangParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tlangParserPLUS || _la == tlangParserMINUS {
			{
				p.SetState(483)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Integer_literalContext).sign = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == tlangParserPLUS || _la == tlangParserMINUS) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Integer_literalContext).sign = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(486)
			p.Match(tlangParserDECIMAL_LITERAL)
		}

	case tlangParserBINARY_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(487)
			p.Match(tlangParserBINARY_LITERAL)
		}

	case tlangParserHEX_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(488)
			p.Match(tlangParserHEX_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = tlangParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlangParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) BLOCK_COMMENT() antlr.TerminalNode {
	return s.GetToken(tlangParserBLOCK_COMMENT, 0)
}

func (s *CommentContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(tlangParserLINE_COMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlangListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tlangVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tlangParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, tlangParserRULE_comment)
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
		p.SetState(491)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tlangParserBLOCK_COMMENT || _la == tlangParserLINE_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *tlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 28:
		var t *PutItemListContext = nil
		if localctx != nil {
			t = localctx.(*PutItemListContext)
		}
		return p.PutItemList_Sempred(t, predIndex)

	case 39:
		var t *IndexTypeListContext = nil
		if localctx != nil {
			t = localctx.(*IndexTypeListContext)
		}
		return p.IndexTypeList_Sempred(t, predIndex)

	case 47:
		var t *IdentifierListContext = nil
		if localctx != nil {
			t = localctx.(*IdentifierListContext)
		}
		return p.IdentifierList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tlangParser) ExportList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *tlangParser) ParamDeclarationList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *tlangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *tlangParser) PutItemList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *tlangParser) IndexTypeList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *tlangParser) IdentifierList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
