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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 108, 603,
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
	60, 4, 61, 9, 61, 4, 62, 9, 62, 3, 2, 6, 2, 126, 10, 2, 13, 2, 14, 2, 127,
	3, 3, 3, 3, 5, 3, 132, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 138, 10, 4,
	3, 4, 5, 4, 141, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 147, 10, 5, 3, 5,
	5, 5, 150, 10, 5, 3, 5, 5, 5, 153, 10, 5, 3, 5, 3, 5, 3, 5, 3, 6, 5, 6,
	159, 10, 6, 3, 6, 3, 6, 5, 6, 163, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 6, 9, 175, 10, 9, 13, 9, 14, 9, 176, 3, 9,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 186, 10, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 195, 10, 11, 3, 12, 5, 12, 198,
	10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13,
	208, 10, 13, 12, 13, 14, 13, 211, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 220, 10, 14, 3, 15, 3, 15, 5, 15, 224, 10, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 229, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 235, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 241, 10, 16, 3, 17,
	3, 17, 3, 18, 5, 18, 246, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 253, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 261,
	10, 19, 12, 19, 14, 19, 264, 11, 19, 3, 20, 6, 20, 267, 10, 20, 13, 20,
	14, 20, 268, 3, 21, 3, 21, 5, 21, 273, 10, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 5, 22, 281, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 7, 23, 289, 10, 23, 12, 23, 14, 23, 292, 11, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 305,
	10, 26, 3, 26, 3, 26, 3, 26, 5, 26, 310, 10, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 7, 26, 319, 10, 26, 12, 26, 14, 26, 322, 11, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 331, 10, 27, 3,
	28, 3, 28, 3, 28, 5, 28, 336, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28,
	342, 10, 28, 7, 28, 344, 10, 28, 12, 28, 14, 28, 347, 11, 28, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 357, 10, 30, 12, 30,
	14, 30, 360, 11, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 375, 10, 31, 3, 31, 7, 31,
	378, 10, 31, 12, 31, 14, 31, 381, 11, 31, 3, 32, 3, 32, 3, 32, 5, 32, 386,
	10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 394, 10, 33, 12,
	33, 14, 33, 397, 11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34,
	405, 10, 34, 12, 34, 14, 34, 408, 11, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 7, 35, 416, 10, 35, 12, 35, 14, 35, 419, 11, 35, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 429, 10, 37, 12, 37, 14,
	37, 432, 11, 37, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 438, 10, 38, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 450,
	10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 7, 43, 467, 10, 43, 12, 43, 14,
	43, 470, 11, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46,
	3, 46, 3, 46, 5, 46, 482, 10, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 5,
	47, 489, 10, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 50, 3, 50, 5, 50, 513, 10, 50, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 7, 52, 525, 10, 52, 12, 52, 14,
	52, 528, 11, 52, 3, 53, 3, 53, 3, 53, 3, 53, 5, 53, 534, 10, 53, 3, 54,
	3, 54, 6, 54, 538, 10, 54, 13, 54, 14, 54, 539, 3, 54, 3, 54, 3, 54, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56,
	555, 10, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 562, 10, 56, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 7, 57, 570, 10, 57, 12, 57, 14,
	57, 573, 11, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59,
	3, 59, 3, 59, 3, 59, 5, 59, 586, 10, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3,
	60, 3, 60, 7, 60, 594, 10, 60, 12, 60, 14, 60, 597, 11, 60, 3, 61, 3, 61,
	3, 62, 3, 62, 3, 62, 2, 17, 24, 36, 44, 50, 54, 58, 60, 64, 66, 68, 72,
	84, 102, 112, 118, 63, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
	64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98,
	100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 2, 10, 4, 2,
	13, 13, 16, 16, 3, 2, 58, 60, 4, 2, 3, 5, 12, 12, 4, 2, 74, 78, 90, 91,
	4, 2, 72, 73, 92, 92, 5, 2, 71, 71, 93, 97, 102, 103, 3, 2, 6, 7, 3, 2,
	107, 108, 2, 630, 2, 125, 3, 2, 2, 2, 4, 131, 3, 2, 2, 2, 6, 133, 3, 2,
	2, 2, 8, 142, 3, 2, 2, 2, 10, 158, 3, 2, 2, 2, 12, 164, 3, 2, 2, 2, 14,
	169, 3, 2, 2, 2, 16, 171, 3, 2, 2, 2, 18, 185, 3, 2, 2, 2, 20, 194, 3,
	2, 2, 2, 22, 197, 3, 2, 2, 2, 24, 201, 3, 2, 2, 2, 26, 219, 3, 2, 2, 2,
	28, 234, 3, 2, 2, 2, 30, 240, 3, 2, 2, 2, 32, 242, 3, 2, 2, 2, 34, 252,
	3, 2, 2, 2, 36, 254, 3, 2, 2, 2, 38, 266, 3, 2, 2, 2, 40, 272, 3, 2, 2,
	2, 42, 280, 3, 2, 2, 2, 44, 282, 3, 2, 2, 2, 46, 293, 3, 2, 2, 2, 48, 297,
	3, 2, 2, 2, 50, 304, 3, 2, 2, 2, 52, 330, 3, 2, 2, 2, 54, 335, 3, 2, 2,
	2, 56, 348, 3, 2, 2, 2, 58, 350, 3, 2, 2, 2, 60, 361, 3, 2, 2, 2, 62, 385,
	3, 2, 2, 2, 64, 387, 3, 2, 2, 2, 66, 398, 3, 2, 2, 2, 68, 409, 3, 2, 2,
	2, 70, 420, 3, 2, 2, 2, 72, 422, 3, 2, 2, 2, 74, 437, 3, 2, 2, 2, 76, 449,
	3, 2, 2, 2, 78, 451, 3, 2, 2, 2, 80, 455, 3, 2, 2, 2, 82, 458, 3, 2, 2,
	2, 84, 460, 3, 2, 2, 2, 86, 471, 3, 2, 2, 2, 88, 475, 3, 2, 2, 2, 90, 478,
	3, 2, 2, 2, 92, 485, 3, 2, 2, 2, 94, 492, 3, 2, 2, 2, 96, 495, 3, 2, 2,
	2, 98, 512, 3, 2, 2, 2, 100, 514, 3, 2, 2, 2, 102, 518, 3, 2, 2, 2, 104,
	529, 3, 2, 2, 2, 106, 535, 3, 2, 2, 2, 108, 544, 3, 2, 2, 2, 110, 561,
	3, 2, 2, 2, 112, 563, 3, 2, 2, 2, 114, 574, 3, 2, 2, 2, 116, 585, 3, 2,
	2, 2, 118, 587, 3, 2, 2, 2, 120, 598, 3, 2, 2, 2, 122, 600, 3, 2, 2, 2,
	124, 126, 5, 4, 3, 2, 125, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 3, 3, 2, 2, 2, 129, 132, 5,
	40, 21, 2, 130, 132, 5, 16, 9, 2, 131, 129, 3, 2, 2, 2, 131, 130, 3, 2,
	2, 2, 132, 5, 3, 2, 2, 2, 133, 134, 9, 2, 2, 2, 134, 140, 7, 105, 2, 2,
	135, 137, 7, 64, 2, 2, 136, 138, 5, 36, 19, 2, 137, 136, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 141, 7, 65, 2, 2, 140, 135,
	3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 7, 3, 2, 2, 2, 142, 143, 7, 14,
	2, 2, 143, 149, 7, 105, 2, 2, 144, 146, 7, 64, 2, 2, 145, 147, 5, 36, 19,
	2, 146, 145, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148,
	150, 7, 65, 2, 2, 149, 144, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152,
	3, 2, 2, 2, 151, 153, 7, 105, 2, 2, 152, 151, 3, 2, 2, 2, 152, 153, 3,
	2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 7, 63, 2, 2, 155, 156, 5, 98,
	50, 2, 156, 9, 3, 2, 2, 2, 157, 159, 5, 14, 8, 2, 158, 157, 3, 2, 2, 2,
	158, 159, 3, 2, 2, 2, 159, 162, 3, 2, 2, 2, 160, 163, 5, 8, 5, 2, 161,
	163, 5, 6, 4, 2, 162, 160, 3, 2, 2, 2, 162, 161, 3, 2, 2, 2, 163, 11, 3,
	2, 2, 2, 164, 165, 5, 10, 6, 2, 165, 166, 5, 38, 20, 2, 166, 167, 7, 8,
	2, 2, 167, 168, 7, 105, 2, 2, 168, 13, 3, 2, 2, 2, 169, 170, 9, 3, 2, 2,
	170, 15, 3, 2, 2, 2, 171, 172, 7, 15, 2, 2, 172, 174, 7, 105, 2, 2, 173,
	175, 5, 18, 10, 2, 174, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 174,
	3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 7, 8,
	2, 2, 179, 180, 7, 105, 2, 2, 180, 17, 3, 2, 2, 2, 181, 186, 5, 20, 11,
	2, 182, 186, 5, 26, 14, 2, 183, 186, 5, 28, 15, 2, 184, 186, 5, 40, 21,
	2, 185, 181, 3, 2, 2, 2, 185, 182, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185,
	184, 3, 2, 2, 2, 186, 19, 3, 2, 2, 2, 187, 188, 7, 38, 2, 2, 188, 195,
	5, 24, 13, 2, 189, 190, 7, 38, 2, 2, 190, 191, 7, 64, 2, 2, 191, 192, 5,
	24, 13, 2, 192, 193, 7, 65, 2, 2, 193, 195, 3, 2, 2, 2, 194, 187, 3, 2,
	2, 2, 194, 189, 3, 2, 2, 2, 195, 21, 3, 2, 2, 2, 196, 198, 5, 32, 17, 2,
	197, 196, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199,
	200, 7, 105, 2, 2, 200, 23, 3, 2, 2, 2, 201, 202, 8, 13, 1, 2, 202, 203,
	5, 22, 12, 2, 203, 209, 3, 2, 2, 2, 204, 205, 12, 3, 2, 2, 205, 206, 7,
	68, 2, 2, 206, 208, 5, 22, 12, 2, 207, 204, 3, 2, 2, 2, 208, 211, 3, 2,
	2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 25, 3, 2, 2, 2,
	211, 209, 3, 2, 2, 2, 212, 213, 7, 40, 2, 2, 213, 220, 5, 30, 16, 2, 214,
	215, 7, 40, 2, 2, 215, 216, 7, 64, 2, 2, 216, 217, 5, 30, 16, 2, 217, 218,
	7, 65, 2, 2, 218, 220, 3, 2, 2, 2, 219, 212, 3, 2, 2, 2, 219, 214, 3, 2,
	2, 2, 220, 27, 3, 2, 2, 2, 221, 223, 7, 41, 2, 2, 222, 224, 7, 42, 2, 2,
	223, 222, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225,
	235, 5, 30, 16, 2, 226, 228, 7, 41, 2, 2, 227, 229, 7, 42, 2, 2, 228, 227,
	3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 7, 64,
	2, 2, 231, 232, 5, 30, 16, 2, 232, 233, 7, 65, 2, 2, 233, 235, 3, 2, 2,
	2, 234, 221, 3, 2, 2, 2, 234, 226, 3, 2, 2, 2, 235, 29, 3, 2, 2, 2, 236,
	241, 7, 105, 2, 2, 237, 238, 7, 105, 2, 2, 238, 239, 7, 89, 2, 2, 239,
	241, 7, 6, 2, 2, 240, 236, 3, 2, 2, 2, 240, 237, 3, 2, 2, 2, 241, 31, 3,
	2, 2, 2, 242, 243, 9, 4, 2, 2, 243, 33, 3, 2, 2, 2, 244, 246, 7, 12, 2,
	2, 245, 244, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247,
	248, 5, 118, 60, 2, 248, 249, 7, 63, 2, 2, 249, 250, 5, 98, 50, 2, 250,
	253, 3, 2, 2, 2, 251, 253, 5, 10, 6, 2, 252, 245, 3, 2, 2, 2, 252, 251,
	3, 2, 2, 2, 253, 35, 3, 2, 2, 2, 254, 255, 8, 19, 1, 2, 255, 256, 5, 34,
	18, 2, 256, 262, 3, 2, 2, 2, 257, 258, 12, 3, 2, 2, 258, 259, 7, 68, 2,
	2, 259, 261, 5, 34, 18, 2, 260, 257, 3, 2, 2, 2, 261, 264, 3, 2, 2, 2,
	262, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 37, 3, 2, 2, 2, 264, 262,
	3, 2, 2, 2, 265, 267, 5, 40, 21, 2, 266, 265, 3, 2, 2, 2, 267, 268, 3,
	2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 39, 3, 2, 2,
	2, 270, 273, 5, 76, 39, 2, 271, 273, 5, 74, 38, 2, 272, 270, 3, 2, 2, 2,
	272, 271, 3, 2, 2, 2, 273, 41, 3, 2, 2, 2, 274, 281, 5, 120, 61, 2, 275,
	281, 7, 105, 2, 2, 276, 277, 7, 64, 2, 2, 277, 278, 5, 70, 36, 2, 278,
	279, 7, 65, 2, 2, 279, 281, 3, 2, 2, 2, 280, 274, 3, 2, 2, 2, 280, 275,
	3, 2, 2, 2, 280, 276, 3, 2, 2, 2, 281, 43, 3, 2, 2, 2, 282, 283, 8, 23,
	1, 2, 283, 284, 5, 70, 36, 2, 284, 290, 3, 2, 2, 2, 285, 286, 12, 3, 2,
	2, 286, 287, 7, 68, 2, 2, 287, 289, 5, 44, 23, 4, 288, 285, 3, 2, 2, 2,
	289, 292, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291,
	45, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 293, 294, 5, 42, 22, 2, 294, 295,
	7, 79, 2, 2, 295, 296, 5, 42, 22, 2, 296, 47, 3, 2, 2, 2, 297, 298, 7,
	62, 2, 2, 298, 299, 5, 42, 22, 2, 299, 49, 3, 2, 2, 2, 300, 301, 8, 26,
	1, 2, 301, 305, 5, 42, 22, 2, 302, 305, 5, 48, 25, 2, 303, 305, 5, 46,
	24, 2, 304, 300, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 303, 3, 2, 2, 2,
	305, 320, 3, 2, 2, 2, 306, 307, 12, 5, 2, 2, 307, 309, 7, 64, 2, 2, 308,
	310, 5, 44, 23, 2, 309, 308, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 311,
	3, 2, 2, 2, 311, 319, 7, 65, 2, 2, 312, 313, 12, 4, 2, 2, 313, 314, 7,
	66, 2, 2, 314, 319, 7, 105, 2, 2, 315, 316, 12, 3, 2, 2, 316, 317, 7, 70,
	2, 2, 317, 319, 7, 105, 2, 2, 318, 306, 3, 2, 2, 2, 318, 312, 3, 2, 2,
	2, 318, 315, 3, 2, 2, 2, 319, 322, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320,
	321, 3, 2, 2, 2, 321, 51, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 323, 331, 5,
	50, 26, 2, 324, 325, 7, 72, 2, 2, 325, 331, 5, 52, 27, 2, 326, 327, 7,
	73, 2, 2, 327, 331, 5, 52, 27, 2, 328, 329, 7, 69, 2, 2, 329, 331, 5, 52,
	27, 2, 330, 323, 3, 2, 2, 2, 330, 324, 3, 2, 2, 2, 330, 326, 3, 2, 2, 2,
	330, 328, 3, 2, 2, 2, 331, 53, 3, 2, 2, 2, 332, 333, 8, 28, 1, 2, 333,
	336, 5, 50, 26, 2, 334, 336, 5, 52, 27, 2, 335, 332, 3, 2, 2, 2, 335, 334,
	3, 2, 2, 2, 336, 345, 3, 2, 2, 2, 337, 338, 12, 3, 2, 2, 338, 341, 5, 56,
	29, 2, 339, 342, 5, 50, 26, 2, 340, 342, 5, 52, 27, 2, 341, 339, 3, 2,
	2, 2, 341, 340, 3, 2, 2, 2, 342, 344, 3, 2, 2, 2, 343, 337, 3, 2, 2, 2,
	344, 347, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346,
	55, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 348, 349, 9, 5, 2, 2, 349, 57, 3,
	2, 2, 2, 350, 351, 8, 30, 1, 2, 351, 352, 5, 54, 28, 2, 352, 358, 3, 2,
	2, 2, 353, 354, 12, 3, 2, 2, 354, 355, 9, 6, 2, 2, 355, 357, 5, 54, 28,
	2, 356, 353, 3, 2, 2, 2, 357, 360, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358,
	359, 3, 2, 2, 2, 359, 59, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 361, 362, 8,
	31, 1, 2, 362, 363, 5, 58, 30, 2, 363, 379, 3, 2, 2, 2, 364, 374, 12, 3,
	2, 2, 365, 375, 7, 80, 2, 2, 366, 375, 7, 81, 2, 2, 367, 375, 7, 82, 2,
	2, 368, 375, 7, 83, 2, 2, 369, 375, 7, 84, 2, 2, 370, 375, 7, 85, 2, 2,
	371, 375, 7, 89, 2, 2, 372, 373, 7, 61, 2, 2, 373, 375, 7, 89, 2, 2, 374,
	365, 3, 2, 2, 2, 374, 366, 3, 2, 2, 2, 374, 367, 3, 2, 2, 2, 374, 368,
	3, 2, 2, 2, 374, 369, 3, 2, 2, 2, 374, 370, 3, 2, 2, 2, 374, 371, 3, 2,
	2, 2, 374, 372, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 378, 5, 58, 30,
	2, 377, 364, 3, 2, 2, 2, 378, 381, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 379,
	380, 3, 2, 2, 2, 380, 61, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 382, 386, 5,
	60, 31, 2, 383, 384, 7, 61, 2, 2, 384, 386, 5, 62, 32, 2, 385, 382, 3,
	2, 2, 2, 385, 383, 3, 2, 2, 2, 386, 63, 3, 2, 2, 2, 387, 388, 8, 33, 1,
	2, 388, 389, 5, 62, 32, 2, 389, 395, 3, 2, 2, 2, 390, 391, 12, 3, 2, 2,
	391, 392, 7, 86, 2, 2, 392, 394, 5, 62, 32, 2, 393, 390, 3, 2, 2, 2, 394,
	397, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 65, 3,
	2, 2, 2, 397, 395, 3, 2, 2, 2, 398, 399, 8, 34, 1, 2, 399, 400, 5, 64,
	33, 2, 400, 406, 3, 2, 2, 2, 401, 402, 12, 3, 2, 2, 402, 403, 7, 87, 2,
	2, 403, 405, 5, 64, 33, 2, 404, 401, 3, 2, 2, 2, 405, 408, 3, 2, 2, 2,
	406, 404, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 67, 3, 2, 2, 2, 408, 406,
	3, 2, 2, 2, 409, 410, 8, 35, 1, 2, 410, 411, 5, 66, 34, 2, 411, 417, 3,
	2, 2, 2, 412, 413, 12, 3, 2, 2, 413, 414, 7, 88, 2, 2, 414, 416, 5, 66,
	34, 2, 415, 412, 3, 2, 2, 2, 416, 419, 3, 2, 2, 2, 417, 415, 3, 2, 2, 2,
	417, 418, 3, 2, 2, 2, 418, 69, 3, 2, 2, 2, 419, 417, 3, 2, 2, 2, 420, 421,
	5, 68, 35, 2, 421, 71, 3, 2, 2, 2, 422, 423, 8, 37, 1, 2, 423, 424, 5,
	70, 36, 2, 424, 430, 3, 2, 2, 2, 425, 426, 12, 3, 2, 2, 426, 427, 7, 68,
	2, 2, 427, 429, 5, 70, 36, 2, 428, 425, 3, 2, 2, 2, 429, 432, 3, 2, 2,
	2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 73, 3, 2, 2, 2, 432,
	430, 3, 2, 2, 2, 433, 438, 5, 96, 49, 2, 434, 438, 5, 110, 56, 2, 435,
	438, 5, 116, 59, 2, 436, 438, 5, 12, 7, 2, 437, 433, 3, 2, 2, 2, 437, 434,
	3, 2, 2, 2, 437, 435, 3, 2, 2, 2, 437, 436, 3, 2, 2, 2, 438, 75, 3, 2,
	2, 2, 439, 450, 5, 70, 36, 2, 440, 450, 5, 78, 40, 2, 441, 450, 5, 80,
	41, 2, 442, 450, 7, 19, 2, 2, 443, 450, 5, 86, 44, 2, 444, 450, 7, 26,
	2, 2, 445, 450, 5, 88, 45, 2, 446, 450, 5, 90, 46, 2, 447, 450, 5, 92,
	47, 2, 448, 450, 5, 94, 48, 2, 449, 439, 3, 2, 2, 2, 449, 440, 3, 2, 2,
	2, 449, 441, 3, 2, 2, 2, 449, 442, 3, 2, 2, 2, 449, 443, 3, 2, 2, 2, 449,
	444, 3, 2, 2, 2, 449, 445, 3, 2, 2, 2, 449, 446, 3, 2, 2, 2, 449, 447,
	3, 2, 2, 2, 449, 448, 3, 2, 2, 2, 450, 77, 3, 2, 2, 2, 451, 452, 5, 50,
	26, 2, 452, 453, 9, 7, 2, 2, 453, 454, 5, 70, 36, 2, 454, 79, 3, 2, 2,
	2, 455, 456, 7, 43, 2, 2, 456, 457, 5, 84, 43, 2, 457, 81, 3, 2, 2, 2,
	458, 459, 5, 76, 39, 2, 459, 83, 3, 2, 2, 2, 460, 461, 8, 43, 1, 2, 461,
	462, 5, 82, 42, 2, 462, 468, 3, 2, 2, 2, 463, 464, 12, 3, 2, 2, 464, 465,
	7, 68, 2, 2, 465, 467, 5, 82, 42, 2, 466, 463, 3, 2, 2, 2, 467, 470, 3,
	2, 2, 2, 468, 466, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 85, 3, 2, 2,
	2, 470, 468, 3, 2, 2, 2, 471, 472, 7, 25, 2, 2, 472, 473, 5, 40, 21, 2,
	473, 474, 7, 8, 2, 2, 474, 87, 3, 2, 2, 2, 475, 476, 7, 27, 2, 2, 476,
	477, 5, 70, 36, 2, 477, 89, 3, 2, 2, 2, 478, 481, 7, 28, 2, 2, 479, 480,
	7, 105, 2, 2, 480, 482, 7, 68, 2, 2, 481, 479, 3, 2, 2, 2, 481, 482, 3,
	2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 484, 7, 105, 2, 2, 484, 91, 3, 2, 2,
	2, 485, 488, 7, 29, 2, 2, 486, 487, 7, 105, 2, 2, 487, 489, 7, 68, 2, 2,
	488, 486, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490,
	491, 7, 105, 2, 2, 491, 93, 3, 2, 2, 2, 492, 493, 7, 31, 2, 2, 493, 494,
	5, 70, 36, 2, 494, 95, 3, 2, 2, 2, 495, 496, 7, 11, 2, 2, 496, 497, 7,
	105, 2, 2, 497, 498, 7, 63, 2, 2, 498, 499, 5, 98, 50, 2, 499, 97, 3, 2,
	2, 2, 500, 513, 7, 44, 2, 2, 501, 513, 7, 45, 2, 2, 502, 513, 7, 46, 2,
	2, 503, 513, 7, 53, 2, 2, 504, 513, 7, 54, 2, 2, 505, 513, 7, 55, 2, 2,
	506, 513, 7, 56, 2, 2, 507, 513, 7, 57, 2, 2, 508, 513, 5, 100, 51, 2,
	509, 513, 5, 104, 53, 2, 510, 513, 5, 106, 54, 2, 511, 513, 7, 105, 2,
	2, 512, 500, 3, 2, 2, 2, 512, 501, 3, 2, 2, 2, 512, 502, 3, 2, 2, 2, 512,
	503, 3, 2, 2, 2, 512, 504, 3, 2, 2, 2, 512, 505, 3, 2, 2, 2, 512, 506,
	3, 2, 2, 2, 512, 507, 3, 2, 2, 2, 512, 508, 3, 2, 2, 2, 512, 509, 3, 2,
	2, 2, 512, 510, 3, 2, 2, 2, 512, 511, 3, 2, 2, 2, 513, 99, 3, 2, 2, 2,
	514, 515, 7, 7, 2, 2, 515, 516, 7, 67, 2, 2, 516, 517, 7, 7, 2, 2, 517,
	101, 3, 2, 2, 2, 518, 519, 8, 52, 1, 2, 519, 520, 5, 100, 51, 2, 520, 526,
	3, 2, 2, 2, 521, 522, 12, 3, 2, 2, 522, 523, 7, 68, 2, 2, 523, 525, 5,
	100, 51, 2, 524, 521, 3, 2, 2, 2, 525, 528, 3, 2, 2, 2, 526, 524, 3, 2,
	2, 2, 526, 527, 3, 2, 2, 2, 527, 103, 3, 2, 2, 2, 528, 526, 3, 2, 2, 2,
	529, 533, 7, 47, 2, 2, 530, 531, 7, 64, 2, 2, 531, 532, 7, 7, 2, 2, 532,
	534, 7, 65, 2, 2, 533, 530, 3, 2, 2, 2, 533, 534, 3, 2, 2, 2, 534, 105,
	3, 2, 2, 2, 535, 537, 7, 50, 2, 2, 536, 538, 5, 108, 55, 2, 537, 536, 3,
	2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 537, 3, 2, 2, 2, 539, 540, 3, 2, 2,
	2, 540, 541, 3, 2, 2, 2, 541, 542, 7, 8, 2, 2, 542, 543, 7, 50, 2, 2, 543,
	107, 3, 2, 2, 2, 544, 545, 5, 118, 60, 2, 545, 546, 7, 63, 2, 2, 546, 547,
	5, 98, 50, 2, 547, 109, 3, 2, 2, 2, 548, 549, 7, 12, 2, 2, 549, 550, 5,
	112, 57, 2, 550, 551, 7, 63, 2, 2, 551, 554, 5, 98, 50, 2, 552, 553, 7,
	71, 2, 2, 553, 555, 5, 70, 36, 2, 554, 552, 3, 2, 2, 2, 554, 555, 3, 2,
	2, 2, 555, 562, 3, 2, 2, 2, 556, 557, 7, 12, 2, 2, 557, 558, 5, 112, 57,
	2, 558, 559, 7, 71, 2, 2, 559, 560, 5, 70, 36, 2, 560, 562, 3, 2, 2, 2,
	561, 548, 3, 2, 2, 2, 561, 556, 3, 2, 2, 2, 562, 111, 3, 2, 2, 2, 563,
	564, 8, 57, 1, 2, 564, 565, 5, 114, 58, 2, 565, 571, 3, 2, 2, 2, 566, 567,
	12, 3, 2, 2, 567, 568, 7, 68, 2, 2, 568, 570, 5, 114, 58, 2, 569, 566,
	3, 2, 2, 2, 570, 573, 3, 2, 2, 2, 571, 569, 3, 2, 2, 2, 571, 572, 3, 2,
	2, 2, 572, 113, 3, 2, 2, 2, 573, 571, 3, 2, 2, 2, 574, 575, 7, 105, 2,
	2, 575, 115, 3, 2, 2, 2, 576, 577, 7, 48, 2, 2, 577, 578, 5, 100, 51, 2,
	578, 579, 7, 9, 2, 2, 579, 580, 5, 98, 50, 2, 580, 586, 3, 2, 2, 2, 581,
	582, 7, 48, 2, 2, 582, 583, 5, 100, 51, 2, 583, 584, 7, 68, 2, 2, 584,
	586, 3, 2, 2, 2, 585, 576, 3, 2, 2, 2, 585, 581, 3, 2, 2, 2, 586, 117,
	3, 2, 2, 2, 587, 588, 8, 60, 1, 2, 588, 589, 7, 105, 2, 2, 589, 595, 3,
	2, 2, 2, 590, 591, 12, 3, 2, 2, 591, 592, 7, 68, 2, 2, 592, 594, 7, 105,
	2, 2, 593, 590, 3, 2, 2, 2, 594, 597, 3, 2, 2, 2, 595, 593, 3, 2, 2, 2,
	595, 596, 3, 2, 2, 2, 596, 119, 3, 2, 2, 2, 597, 595, 3, 2, 2, 2, 598,
	599, 9, 8, 2, 2, 599, 121, 3, 2, 2, 2, 600, 601, 9, 9, 2, 2, 601, 123,
	3, 2, 2, 2, 58, 127, 131, 137, 140, 146, 149, 152, 158, 162, 176, 185,
	194, 197, 209, 219, 223, 228, 234, 240, 245, 252, 262, 268, 272, 280, 290,
	304, 309, 318, 320, 330, 335, 341, 345, 358, 374, 379, 385, 395, 406, 417,
	430, 437, 449, 468, 481, 488, 512, 526, 533, 539, 554, 561, 571, 585, 595,
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
	"recordType", "recordField", "variableDeclaration", "variableIdentifierList",
	"variableIdentifier", "arrayDeclaration", "identifierList", "literal",
	"comment",
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
	TuringParserRULE_variableIdentifierList   = 55
	TuringParserRULE_variableIdentifier       = 56
	TuringParserRULE_arrayDeclaration         = 57
	TuringParserRULE_identifierList           = 58
	TuringParserRULE_literal                  = 59
	TuringParserRULE_comment                  = 60
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
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserCLASS)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(TuringParserPUT-41))|(1<<(TuringParserARRAY-41))|(1<<(TuringParserDEFERRED-41))|(1<<(TuringParserFORWARD-41))|(1<<(TuringParserBODY-41))|(1<<(TuringParserNOT-41))|(1<<(TuringParserCARET-41))|(1<<(TuringParserL_PAREN-41))|(1<<(TuringParserCHEAT-41))|(1<<(TuringParserPLUS-41))|(1<<(TuringParserMINUS-41)))) != 0) || _la == TuringParserIDENTIFIER {
		{
			p.SetState(122)
			p.TopLevel()
		}

		p.SetState(125)
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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.StatementOrDeclaration()
		}

	case TuringParserCLASS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
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
		p.SetState(131)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TuringParserPROCEDURE || _la == TuringParserPROCESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(132)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(133)
			p.Match(TuringParserL_PAREN)
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS))) != 0) || (((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(TuringParserDEFERRED-56))|(1<<(TuringParserFORWARD-56))|(1<<(TuringParserBODY-56)))) != 0) || _la == TuringParserIDENTIFIER {
			{
				p.SetState(134)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(137)
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
		p.SetState(140)
		p.Match(TuringParserFUNCTION)
	}
	{
		p.SetState(141)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TuringParserL_PAREN {
		{
			p.SetState(142)
			p.Match(TuringParserL_PAREN)
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS))) != 0) || (((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(TuringParserDEFERRED-56))|(1<<(TuringParserFORWARD-56))|(1<<(TuringParserBODY-56)))) != 0) || _la == TuringParserIDENTIFIER {
			{
				p.SetState(143)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(146)
			p.Match(TuringParserR_PAREN)
		}

	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TuringParserIDENTIFIER {
		{
			p.SetState(149)
			p.Match(TuringParserIDENTIFIER)
		}

	}
	{
		p.SetState(152)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(153)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(TuringParserDEFERRED-56))|(1<<(TuringParserFORWARD-56))|(1<<(TuringParserBODY-56)))) != 0 {
		{
			p.SetState(155)
			p.SubprogramPrefix()
		}

	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserFUNCTION:
		{
			p.SetState(158)
			p.FuncHeader()
		}

	case TuringParserPROCEDURE, TuringParserPROCESS:
		{
			p.SetState(159)
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
		p.SetState(162)
		p.SubprogramHeader()
	}
	{
		p.SetState(163)
		p.ProcBody()
	}
	{
		p.SetState(164)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(165)
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
		p.SetState(167)
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
		p.SetState(169)
		p.Match(TuringParserCLASS)
	}
	{
		p.SetState(170)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(TuringParserEXPORT-36))|(1<<(TuringParserINHERIT-36))|(1<<(TuringParserIMPLEMENT-36))|(1<<(TuringParserPUT-36))|(1<<(TuringParserARRAY-36))|(1<<(TuringParserDEFERRED-36))|(1<<(TuringParserFORWARD-36))|(1<<(TuringParserBODY-36))|(1<<(TuringParserNOT-36))|(1<<(TuringParserCARET-36))|(1<<(TuringParserL_PAREN-36))|(1<<(TuringParserCHEAT-36)))) != 0) || _la == TuringParserPLUS || _la == TuringParserMINUS || _la == TuringParserIDENTIFIER {
		{
			p.SetState(171)
			p.ClassBody()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(177)
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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserEXPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.ExportStatement()
		}

	case TuringParserINHERIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.InheritStatement()
		}

	case TuringParserIMPLEMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.ImplementStatement()
		}

	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)
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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(TuringParserEXPORT)
		}
		{
			p.SetState(186)
			p.exportList(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Match(TuringParserEXPORT)
		}
		{
			p.SetState(188)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(189)
			p.exportList(0)
		}
		{
			p.SetState(190)
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
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserT__0)|(1<<TuringParserT__1)|(1<<TuringParserT__2)|(1<<TuringParserVAR))) != 0 {
		{
			p.SetState(194)
			p.HowExport()
		}

	}
	{
		p.SetState(197)
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
		p.SetState(200)
		p.ExportListItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(207)
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
			p.SetState(202)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(203)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(204)
				p.ExportListItem()
			}

		}
		p.SetState(209)
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

	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(TuringParserINHERIT)
		}
		{
			p.SetState(211)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(TuringParserINHERIT)
		}
		{
			p.SetState(213)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(214)
			p.IdOrFileItem()
		}
		{
			p.SetState(215)
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

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Match(TuringParserIMPLEMENT)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserBY {
			{
				p.SetState(220)
				p.Match(TuringParserBY)
			}

		}
		{
			p.SetState(223)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(TuringParserIMPLEMENT)
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserBY {
			{
				p.SetState(225)
				p.Match(TuringParserBY)
			}

		}
		{
			p.SetState(228)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(229)
			p.IdOrFileItem()
		}
		{
			p.SetState(230)
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

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(TuringParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(236)
			p.Match(TuringParserIN)
		}
		{
			p.SetState(237)
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
		p.SetState(240)
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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserVAR, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserVAR {
			{
				p.SetState(242)
				p.Match(TuringParserVAR)
			}

		}
		{
			p.SetState(245)
			p.identifierList(0)
		}
		{
			p.SetState(246)
			p.Match(TuringParserCOLON)
		}
		{
			p.SetState(247)
			p.TypeSpec()
		}

	case TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
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
		p.SetState(253)
		p.ParamDeclaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(260)
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
			p.SetState(255)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(256)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(257)
				p.ParamDeclaration()
			}

		}
		p.SetState(262)
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
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(TuringParserPUT-41))|(1<<(TuringParserARRAY-41))|(1<<(TuringParserDEFERRED-41))|(1<<(TuringParserFORWARD-41))|(1<<(TuringParserBODY-41))|(1<<(TuringParserNOT-41))|(1<<(TuringParserCARET-41))|(1<<(TuringParserL_PAREN-41))|(1<<(TuringParserCHEAT-41))|(1<<(TuringParserPLUS-41))|(1<<(TuringParserMINUS-41)))) != 0) || _la == TuringParserIDENTIFIER {
		{
			p.SetState(263)
			p.StatementOrDeclaration()
		}

		p.SetState(266)
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

	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Statements()
		}

	case TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Literal()
		}

	case TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.Match(TuringParserIDENTIFIER)
		}

	case TuringParserL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(274)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(275)
			p.Expression()
		}
		{
			p.SetState(276)
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
		p.SetState(281)
		p.Expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(288)
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
			p.SetState(283)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(284)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(285)
				p.argumentList(2)
			}

		}
		p.SetState(290)
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
		p.SetState(291)
		p.PrimaryExpression()
	}
	{
		p.SetState(292)
		p.Match(TuringParserEXP)
	}
	{
		p.SetState(293)
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
		p.SetState(295)
		p.Match(TuringParserCARET)
	}
	{
		p.SetState(296)
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
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(299)
			p.PrimaryExpression()
		}

	case 2:
		{
			p.SetState(300)
			p.PointerExpression()
		}

	case 3:
		{
			p.SetState(301)
			p.ExponentialExpression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_postfixExpression)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(305)
					p.Match(TuringParserL_PAREN)
				}
				p.SetState(307)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == TuringParserSTRING_LITERAL || _la == TuringParserINTEGER_LITERAL || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(TuringParserNOT-59))|(1<<(TuringParserCARET-59))|(1<<(TuringParserL_PAREN-59))|(1<<(TuringParserCHEAT-59))|(1<<(TuringParserPLUS-59))|(1<<(TuringParserMINUS-59)))) != 0) || _la == TuringParserIDENTIFIER {
					{
						p.SetState(306)
						p.argumentList(0)
					}

				}
				{
					p.SetState(309)
					p.Match(TuringParserR_PAREN)
				}

			case 2:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_postfixExpression)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(311)
					p.Match(TuringParserDOT)
				}
				{
					p.SetState(312)
					p.Match(TuringParserIDENTIFIER)
				}

			case 3:
				localctx = NewPostfixExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_postfixExpression)
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(314)
					p.Match(TuringParserDEREFERENCE)
				}
				{
					p.SetState(315)
					p.Match(TuringParserIDENTIFIER)
				}

			}

		}
		p.SetState(320)
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

	p.SetState(328)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserCARET, TuringParserL_PAREN, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.postfixExpression(0)
		}

	case TuringParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(TuringParserPLUS)
		}
		{
			p.SetState(323)
			p.PrefixExpression()
		}

	case TuringParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.Match(TuringParserMINUS)
		}
		{
			p.SetState(325)
			p.PrefixExpression()
		}

	case TuringParserCHEAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(326)
			p.Match(TuringParserCHEAT)
		}
		{
			p.SetState(327)
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
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(331)
			p.postfixExpression(0)
		}

	case 2:
		{
			p.SetState(332)
			p.PrefixExpression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(343)
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
			p.SetState(335)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(336)
				p.MultiplicativeOperator()
			}
			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(337)
					p.postfixExpression(0)
				}

			case 2:
				{
					p.SetState(338)
					p.PrefixExpression()
				}

			}

		}
		p.SetState(345)
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
		p.SetState(346)
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
		p.SetState(349)
		p.multiplicativeExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(356)
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
			p.SetState(351)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(352)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(TuringParserPLUS-70))|(1<<(TuringParserMINUS-70))|(1<<(TuringParserXOR-70)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(353)
				p.multiplicativeExpression(0)
			}

		}
		p.SetState(358)
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
		p.SetState(360)
		p.additiveExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(377)
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
			p.SetState(362)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(372)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case TuringParserLESSTHAN:
				{
					p.SetState(363)
					p.Match(TuringParserLESSTHAN)
				}

			case TuringParserGREATERTHAN:
				{
					p.SetState(364)
					p.Match(TuringParserGREATERTHAN)
				}

			case TuringParserEQUAL:
				{
					p.SetState(365)
					p.Match(TuringParserEQUAL)
				}

			case TuringParserLESSTHANEQUAL:
				{
					p.SetState(366)
					p.Match(TuringParserLESSTHANEQUAL)
				}

			case TuringParserGREATERTHANEQUAL:
				{
					p.SetState(367)
					p.Match(TuringParserGREATERTHANEQUAL)
				}

			case TuringParserNOTEQUAL:
				{
					p.SetState(368)
					p.Match(TuringParserNOTEQUAL)
				}

			case TuringParserIN:
				{
					p.SetState(369)
					p.Match(TuringParserIN)
				}

			case TuringParserNOT:
				{
					p.SetState(370)
					p.Match(TuringParserNOT)
				}
				{
					p.SetState(371)
					p.Match(TuringParserIN)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(374)
				p.additiveExpression(0)
			}

		}
		p.SetState(379)
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

	p.SetState(383)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.comparativeExpression(0)
		}

	case TuringParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.Match(TuringParserNOT)
		}
		{
			p.SetState(382)
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
		p.SetState(386)
		p.NotExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(393)
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
			p.SetState(388)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(389)
				p.Match(TuringParserAND)
			}
			{
				p.SetState(390)
				p.NotExpression()
			}

		}
		p.SetState(395)
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
		p.SetState(397)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(404)
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
			p.SetState(399)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(400)
				p.Match(TuringParserOR)
			}
			{
				p.SetState(401)
				p.andExpression(0)
			}

		}
		p.SetState(406)
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
		p.SetState(408)
		p.orExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(415)
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
			p.SetState(410)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(411)
				p.Match(TuringParserIMPLIES)
			}
			{
				p.SetState(412)
				p.orExpression(0)
			}

		}
		p.SetState(417)
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
		p.SetState(418)
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
		p.SetState(421)
		p.Expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(428)
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
			p.SetState(423)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(424)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(425)
				p.Expression()
			}

		}
		p.SetState(430)
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

	p.SetState(435)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.TypeDeclaration()
		}

	case TuringParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(432)
			p.VariableDeclaration()
		}

	case TuringParserARRAY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(433)
			p.ArrayDeclaration()
		}

	case TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(434)
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

	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.AssignmentStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(439)
			p.PutStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(440)
			p.Match(TuringParserEXIT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(441)
			p.BeginStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(442)
			p.Match(TuringParserRETURN)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(443)
			p.ResultStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(444)
			p.NewStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(445)
			p.FreeStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(446)
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
		p.SetState(449)
		p.postfixExpression(0)
	}
	{
		p.SetState(450)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(TuringParserASSIGNMENT-69))|(1<<(TuringParserPLUSEQUALS-69))|(1<<(TuringParserMINUSEQUALS-69))|(1<<(TuringParserMULTIPLYEQUALS-69))|(1<<(TuringParserDIVIDEEQUALS-69))|(1<<(TuringParserDIVEQUALS-69))|(1<<(TuringParserSHREQUALS-69)))) != 0) || _la == TuringParserSHLEQUALS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(451)
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
		p.SetState(453)
		p.Match(TuringParserPUT)
	}
	{
		p.SetState(454)
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
		p.SetState(456)
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
		p.SetState(459)
		p.PutItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(466)
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
			p.SetState(461)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(462)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(463)
				p.PutItem()
			}

		}
		p.SetState(468)
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
		p.SetState(469)
		p.Match(TuringParserBEGIN)
	}
	{
		p.SetState(470)
		p.StatementOrDeclaration()
	}
	{
		p.SetState(471)
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
		p.SetState(473)
		p.Match(TuringParserRESULT)
	}
	{
		p.SetState(474)
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
		p.SetState(476)
		p.Match(TuringParserNEW)
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(477)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(478)
			p.Match(TuringParserCOMMA)
		}

	}
	{
		p.SetState(481)
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
		p.SetState(483)
		p.Match(TuringParserFREE)
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(484)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(485)
			p.Match(TuringParserCOMMA)
		}

	}
	{
		p.SetState(488)
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
		p.SetState(490)
		p.Match(TuringParserFORK)
	}
	{
		p.SetState(491)
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
		p.SetState(493)
		p.Match(TuringParserTYPE)
	}
	{
		p.SetState(494)
		p.Match(TuringParserIDENTIFIER)
	}
	{
		p.SetState(495)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(496)
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

	p.SetState(510)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(498)
			p.Match(TuringParserINT)
		}

	case TuringParserREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(499)
			p.Match(TuringParserREAL)
		}

	case TuringParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(500)
			p.Match(TuringParserBOOLEAN)
		}

	case TuringParserNAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(501)
			p.Match(TuringParserNAT)
		}

	case TuringParserINTN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(502)
			p.Match(TuringParserINTN)
		}

	case TuringParserNATN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(503)
			p.Match(TuringParserNATN)
		}

	case TuringParserREALN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(504)
			p.Match(TuringParserREALN)
		}

	case TuringParserCHAR:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(505)
			p.Match(TuringParserCHAR)
		}

	case TuringParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(506)
			p.IndexType()
		}

	case TuringParserSTRING:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(507)
			p.StringType()
		}

	case TuringParserRECORD:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(508)
			p.RecordType()
		}

	case TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(509)
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
		p.SetState(512)
		p.Match(TuringParserINTEGER_LITERAL)
	}
	{
		p.SetState(513)
		p.Match(TuringParserRANGE)
	}
	{
		p.SetState(514)
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
		p.SetState(517)
		p.IndexType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(524)
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
			p.SetState(519)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(520)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(521)
				p.IndexType()
			}

		}
		p.SetState(526)
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
		p.SetState(527)
		p.Match(TuringParserSTRING)
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(528)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(529)
			p.Match(TuringParserINTEGER_LITERAL)
		}
		{
			p.SetState(530)
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
		p.SetState(533)
		p.Match(TuringParserRECORD)
	}
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TuringParserIDENTIFIER {
		{
			p.SetState(534)
			p.RecordField()
		}

		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(539)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(540)
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
		p.SetState(542)
		p.identifierList(0)
	}
	{
		p.SetState(543)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(544)
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

func (s *VariableDeclarationContext) VariableIdentifierList() IVariableIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableIdentifierListContext)
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

	p.SetState(559)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(546)
			p.Match(TuringParserVAR)
		}
		{
			p.SetState(547)
			p.variableIdentifierList(0)
		}
		{
			p.SetState(548)
			p.Match(TuringParserCOLON)
		}
		{
			p.SetState(549)
			p.TypeSpec()
		}
		p.SetState(552)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserASSIGNMENT {
			{
				p.SetState(550)
				p.Match(TuringParserASSIGNMENT)
			}
			{
				p.SetState(551)
				p.Expression()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(554)
			p.Match(TuringParserVAR)
		}
		{
			p.SetState(555)
			p.variableIdentifierList(0)
		}
		{
			p.SetState(556)
			p.Match(TuringParserASSIGNMENT)
		}
		{
			p.SetState(557)
			p.Expression()
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
	p.RuleIndex = TuringParserRULE_variableIdentifierList
	return p
}

func (*VariableIdentifierListContext) IsVariableIdentifierListContext() {}

func NewVariableIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableIdentifierListContext {
	var p = new(VariableIdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_variableIdentifierList

	return p
}

func (s *VariableIdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableIdentifierListContext) VariableIdentifier() IVariableIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableIdentifierContext)
}

func (s *VariableIdentifierListContext) VariableIdentifierList() IVariableIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableIdentifierListContext)
}

func (s *VariableIdentifierListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, 0)
}

func (s *VariableIdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableIdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableIdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterVariableIdentifierList(s)
	}
}

func (s *VariableIdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitVariableIdentifierList(s)
	}
}

func (p *TuringParser) VariableIdentifierList() (localctx IVariableIdentifierListContext) {
	return p.variableIdentifierList(0)
}

func (p *TuringParser) variableIdentifierList(_p int) (localctx IVariableIdentifierListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewVariableIdentifierListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVariableIdentifierListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 110
	p.EnterRecursionRule(localctx, 110, TuringParserRULE_variableIdentifierList, _p)

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
		p.SetState(562)
		p.VariableIdentifier()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewVariableIdentifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_variableIdentifierList)
			p.SetState(564)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(565)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(566)
				p.VariableIdentifier()
			}

		}
		p.SetState(571)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 112, TuringParserRULE_variableIdentifier)

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
		p.SetState(572)
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
	p.EnterRule(localctx, 114, TuringParserRULE_arrayDeclaration)

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

	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(574)
			p.Match(TuringParserARRAY)
		}
		{
			p.SetState(575)
			p.IndexType()
		}
		{
			p.SetState(576)
			p.Match(TuringParserOF)
		}
		{
			p.SetState(577)
			p.TypeSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(579)
			p.Match(TuringParserARRAY)
		}
		{
			p.SetState(580)
			p.IndexType()
		}
		{
			p.SetState(581)
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
	_startState := 116
	p.EnterRecursionRule(localctx, 116, TuringParserRULE_identifierList, _p)

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
		p.SetState(586)
		p.Match(TuringParserIDENTIFIER)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIdentifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_identifierList)
			p.SetState(588)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(589)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(590)
				p.Match(TuringParserIDENTIFIER)
			}

		}
		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 118, TuringParserRULE_literal)
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
		p.SetState(596)
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
	p.EnterRule(localctx, 120, TuringParserRULE_comment)
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
		p.SetState(598)
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

	case 55:
		var t *VariableIdentifierListContext = nil
		if localctx != nil {
			t = localctx.(*VariableIdentifierListContext)
		}
		return p.VariableIdentifierList_Sempred(t, predIndex)

	case 58:
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

func (p *TuringParser) VariableIdentifierList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) IdentifierList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
