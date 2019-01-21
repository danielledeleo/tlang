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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 110, 480,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 3, 2, 6, 2, 104, 10, 2, 13, 2, 14, 2, 105, 3,
	3, 3, 3, 5, 3, 110, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 116, 10, 4, 3,
	4, 5, 4, 119, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 125, 10, 5, 3, 5, 5,
	5, 128, 10, 5, 3, 5, 5, 5, 131, 10, 5, 3, 5, 3, 5, 3, 5, 3, 6, 5, 6, 137,
	10, 6, 3, 6, 3, 6, 5, 6, 141, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 6, 9, 153, 10, 9, 13, 9, 14, 9, 154, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 164, 10, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 173, 10, 11, 3, 12, 5, 12, 176,
	10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13,
	186, 10, 13, 12, 13, 14, 13, 189, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 198, 10, 14, 3, 15, 3, 15, 5, 15, 202, 10, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 207, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 213, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 219, 10, 16, 3, 17,
	3, 17, 3, 18, 5, 18, 224, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 231, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 239,
	10, 19, 12, 19, 14, 19, 242, 11, 19, 3, 20, 6, 20, 245, 10, 20, 13, 20,
	14, 20, 246, 3, 21, 3, 21, 5, 21, 251, 10, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 5, 22, 259, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 5, 23, 269, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 301, 10, 23, 3, 23, 7, 23, 304,
	10, 23, 12, 23, 14, 23, 307, 11, 23, 3, 24, 3, 24, 3, 24, 7, 24, 312, 10,
	24, 12, 24, 14, 24, 315, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5,
	25, 322, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 5, 26, 333, 10, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 348, 10, 30, 12, 30,
	14, 30, 351, 11, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	33, 3, 33, 3, 33, 5, 33, 363, 10, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34,
	5, 34, 370, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 5, 39, 393, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 405, 10, 41, 12, 41, 14,
	41, 408, 11, 41, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 414, 10, 42, 3, 43,
	3, 43, 6, 43, 418, 10, 43, 13, 43, 14, 43, 419, 3, 43, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45,
	435, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 442, 10, 45, 3,
	46, 3, 46, 3, 46, 7, 46, 447, 10, 46, 12, 46, 14, 46, 450, 11, 46, 3, 47,
	3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 5,
	48, 463, 10, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49, 471,
	10, 49, 12, 49, 14, 49, 474, 11, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51,
	2, 8, 24, 36, 44, 58, 80, 96, 52, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
	96, 98, 100, 2, 14, 4, 2, 15, 15, 18, 18, 3, 2, 60, 62, 4, 2, 3, 5, 14,
	14, 4, 2, 71, 71, 74, 75, 4, 2, 76, 80, 92, 93, 4, 2, 74, 75, 94, 94, 5,
	2, 6, 6, 82, 87, 91, 91, 5, 2, 73, 73, 95, 99, 104, 105, 4, 2, 68, 68,
	72, 72, 4, 2, 46, 48, 55, 59, 3, 2, 7, 9, 3, 2, 109, 110, 2, 501, 2, 103,
	3, 2, 2, 2, 4, 109, 3, 2, 2, 2, 6, 111, 3, 2, 2, 2, 8, 120, 3, 2, 2, 2,
	10, 136, 3, 2, 2, 2, 12, 142, 3, 2, 2, 2, 14, 147, 3, 2, 2, 2, 16, 149,
	3, 2, 2, 2, 18, 163, 3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 175, 3, 2, 2,
	2, 24, 179, 3, 2, 2, 2, 26, 197, 3, 2, 2, 2, 28, 212, 3, 2, 2, 2, 30, 218,
	3, 2, 2, 2, 32, 220, 3, 2, 2, 2, 34, 230, 3, 2, 2, 2, 36, 232, 3, 2, 2,
	2, 38, 244, 3, 2, 2, 2, 40, 250, 3, 2, 2, 2, 42, 258, 3, 2, 2, 2, 44, 268,
	3, 2, 2, 2, 46, 308, 3, 2, 2, 2, 48, 321, 3, 2, 2, 2, 50, 332, 3, 2, 2,
	2, 52, 334, 3, 2, 2, 2, 54, 336, 3, 2, 2, 2, 56, 339, 3, 2, 2, 2, 58, 341,
	3, 2, 2, 2, 60, 352, 3, 2, 2, 2, 62, 356, 3, 2, 2, 2, 64, 359, 3, 2, 2,
	2, 66, 366, 3, 2, 2, 2, 68, 373, 3, 2, 2, 2, 70, 376, 3, 2, 2, 2, 72, 381,
	3, 2, 2, 2, 74, 383, 3, 2, 2, 2, 76, 392, 3, 2, 2, 2, 78, 394, 3, 2, 2,
	2, 80, 398, 3, 2, 2, 2, 82, 409, 3, 2, 2, 2, 84, 415, 3, 2, 2, 2, 86, 424,
	3, 2, 2, 2, 88, 441, 3, 2, 2, 2, 90, 443, 3, 2, 2, 2, 92, 451, 3, 2, 2,
	2, 94, 462, 3, 2, 2, 2, 96, 464, 3, 2, 2, 2, 98, 475, 3, 2, 2, 2, 100,
	477, 3, 2, 2, 2, 102, 104, 5, 4, 3, 2, 103, 102, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 3, 3, 2, 2,
	2, 107, 110, 5, 40, 21, 2, 108, 110, 5, 16, 9, 2, 109, 107, 3, 2, 2, 2,
	109, 108, 3, 2, 2, 2, 110, 5, 3, 2, 2, 2, 111, 112, 9, 2, 2, 2, 112, 118,
	7, 107, 2, 2, 113, 115, 7, 66, 2, 2, 114, 116, 5, 36, 19, 2, 115, 114,
	3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 119, 7, 67,
	2, 2, 118, 113, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 7, 3, 2, 2, 2, 120,
	121, 7, 16, 2, 2, 121, 127, 7, 107, 2, 2, 122, 124, 7, 66, 2, 2, 123, 125,
	5, 36, 19, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126, 3,
	2, 2, 2, 126, 128, 7, 67, 2, 2, 127, 122, 3, 2, 2, 2, 127, 128, 3, 2, 2,
	2, 128, 130, 3, 2, 2, 2, 129, 131, 7, 107, 2, 2, 130, 129, 3, 2, 2, 2,
	130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 7, 65, 2, 2, 133,
	134, 5, 76, 39, 2, 134, 9, 3, 2, 2, 2, 135, 137, 5, 14, 8, 2, 136, 135,
	3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 141, 5, 8,
	5, 2, 139, 141, 5, 6, 4, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2, 2, 2,
	141, 11, 3, 2, 2, 2, 142, 143, 5, 10, 6, 2, 143, 144, 5, 38, 20, 2, 144,
	145, 7, 10, 2, 2, 145, 146, 7, 107, 2, 2, 146, 13, 3, 2, 2, 2, 147, 148,
	9, 3, 2, 2, 148, 15, 3, 2, 2, 2, 149, 150, 7, 17, 2, 2, 150, 152, 7, 107,
	2, 2, 151, 153, 5, 18, 10, 2, 152, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2,
	2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156,
	157, 7, 10, 2, 2, 157, 158, 7, 107, 2, 2, 158, 17, 3, 2, 2, 2, 159, 164,
	5, 20, 11, 2, 160, 164, 5, 26, 14, 2, 161, 164, 5, 28, 15, 2, 162, 164,
	5, 40, 21, 2, 163, 159, 3, 2, 2, 2, 163, 160, 3, 2, 2, 2, 163, 161, 3,
	2, 2, 2, 163, 162, 3, 2, 2, 2, 164, 19, 3, 2, 2, 2, 165, 166, 7, 40, 2,
	2, 166, 173, 5, 24, 13, 2, 167, 168, 7, 40, 2, 2, 168, 169, 7, 66, 2, 2,
	169, 170, 5, 24, 13, 2, 170, 171, 7, 67, 2, 2, 171, 173, 3, 2, 2, 2, 172,
	165, 3, 2, 2, 2, 172, 167, 3, 2, 2, 2, 173, 21, 3, 2, 2, 2, 174, 176, 5,
	32, 17, 2, 175, 174, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 177, 3, 2,
	2, 2, 177, 178, 7, 107, 2, 2, 178, 23, 3, 2, 2, 2, 179, 180, 8, 13, 1,
	2, 180, 181, 5, 22, 12, 2, 181, 187, 3, 2, 2, 2, 182, 183, 12, 3, 2, 2,
	183, 184, 7, 70, 2, 2, 184, 186, 5, 22, 12, 2, 185, 182, 3, 2, 2, 2, 186,
	189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 25, 3,
	2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 191, 7, 42, 2, 2, 191, 198, 5, 30,
	16, 2, 192, 193, 7, 42, 2, 2, 193, 194, 7, 66, 2, 2, 194, 195, 5, 30, 16,
	2, 195, 196, 7, 67, 2, 2, 196, 198, 3, 2, 2, 2, 197, 190, 3, 2, 2, 2, 197,
	192, 3, 2, 2, 2, 198, 27, 3, 2, 2, 2, 199, 201, 7, 43, 2, 2, 200, 202,
	7, 44, 2, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 213, 5, 30, 16, 2, 204, 206, 7, 43, 2, 2, 205, 207, 7, 44, 2,
	2, 206, 205, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208,
	209, 7, 66, 2, 2, 209, 210, 5, 30, 16, 2, 210, 211, 7, 67, 2, 2, 211, 213,
	3, 2, 2, 2, 212, 199, 3, 2, 2, 2, 212, 204, 3, 2, 2, 2, 213, 29, 3, 2,
	2, 2, 214, 219, 7, 107, 2, 2, 215, 216, 7, 107, 2, 2, 216, 217, 7, 91,
	2, 2, 217, 219, 7, 7, 2, 2, 218, 214, 3, 2, 2, 2, 218, 215, 3, 2, 2, 2,
	219, 31, 3, 2, 2, 2, 220, 221, 9, 4, 2, 2, 221, 33, 3, 2, 2, 2, 222, 224,
	7, 14, 2, 2, 223, 222, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 3, 2,
	2, 2, 225, 226, 5, 96, 49, 2, 226, 227, 7, 65, 2, 2, 227, 228, 5, 76, 39,
	2, 228, 231, 3, 2, 2, 2, 229, 231, 5, 10, 6, 2, 230, 223, 3, 2, 2, 2, 230,
	229, 3, 2, 2, 2, 231, 35, 3, 2, 2, 2, 232, 233, 8, 19, 1, 2, 233, 234,
	5, 34, 18, 2, 234, 240, 3, 2, 2, 2, 235, 236, 12, 3, 2, 2, 236, 237, 7,
	70, 2, 2, 237, 239, 5, 34, 18, 2, 238, 235, 3, 2, 2, 2, 239, 242, 3, 2,
	2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 37, 3, 2, 2, 2,
	242, 240, 3, 2, 2, 2, 243, 245, 5, 40, 21, 2, 244, 243, 3, 2, 2, 2, 245,
	246, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 39, 3,
	2, 2, 2, 248, 251, 5, 50, 26, 2, 249, 251, 5, 48, 25, 2, 250, 248, 3, 2,
	2, 2, 250, 249, 3, 2, 2, 2, 251, 41, 3, 2, 2, 2, 252, 259, 5, 98, 50, 2,
	253, 259, 7, 107, 2, 2, 254, 255, 7, 66, 2, 2, 255, 256, 5, 44, 23, 2,
	256, 257, 7, 67, 2, 2, 257, 259, 3, 2, 2, 2, 258, 252, 3, 2, 2, 2, 258,
	253, 3, 2, 2, 2, 258, 254, 3, 2, 2, 2, 259, 43, 3, 2, 2, 2, 260, 261, 8,
	23, 1, 2, 261, 269, 5, 42, 22, 2, 262, 263, 7, 64, 2, 2, 263, 269, 5, 44,
	23, 15, 264, 265, 9, 5, 2, 2, 265, 269, 5, 44, 23, 11, 266, 267, 7, 63,
	2, 2, 267, 269, 5, 44, 23, 7, 268, 260, 3, 2, 2, 2, 268, 262, 3, 2, 2,
	2, 268, 264, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 269, 305, 3, 2, 2, 2, 270,
	271, 12, 13, 2, 2, 271, 272, 7, 81, 2, 2, 272, 304, 5, 44, 23, 14, 273,
	274, 12, 10, 2, 2, 274, 275, 9, 6, 2, 2, 275, 304, 5, 44, 23, 11, 276,
	277, 12, 9, 2, 2, 277, 278, 9, 7, 2, 2, 278, 304, 5, 44, 23, 10, 279, 280,
	12, 8, 2, 2, 280, 281, 9, 8, 2, 2, 281, 304, 5, 44, 23, 9, 282, 283, 12,
	6, 2, 2, 283, 284, 7, 88, 2, 2, 284, 304, 5, 44, 23, 7, 285, 286, 12, 5,
	2, 2, 286, 287, 7, 89, 2, 2, 287, 304, 5, 44, 23, 6, 288, 289, 12, 4, 2,
	2, 289, 290, 7, 90, 2, 2, 290, 304, 5, 44, 23, 5, 291, 292, 12, 3, 2, 2,
	292, 293, 9, 9, 2, 2, 293, 304, 5, 44, 23, 3, 294, 295, 12, 14, 2, 2, 295,
	296, 9, 10, 2, 2, 296, 304, 7, 107, 2, 2, 297, 298, 12, 12, 2, 2, 298,
	300, 7, 66, 2, 2, 299, 301, 5, 46, 24, 2, 300, 299, 3, 2, 2, 2, 300, 301,
	3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 304, 7, 67, 2, 2, 303, 270, 3, 2,
	2, 2, 303, 273, 3, 2, 2, 2, 303, 276, 3, 2, 2, 2, 303, 279, 3, 2, 2, 2,
	303, 282, 3, 2, 2, 2, 303, 285, 3, 2, 2, 2, 303, 288, 3, 2, 2, 2, 303,
	291, 3, 2, 2, 2, 303, 294, 3, 2, 2, 2, 303, 297, 3, 2, 2, 2, 304, 307,
	3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 45, 3, 2,
	2, 2, 307, 305, 3, 2, 2, 2, 308, 313, 5, 44, 23, 2, 309, 310, 7, 70, 2,
	2, 310, 312, 5, 44, 23, 2, 311, 309, 3, 2, 2, 2, 312, 315, 3, 2, 2, 2,
	313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 47, 3, 2, 2, 2, 315, 313,
	3, 2, 2, 2, 316, 322, 5, 70, 36, 2, 317, 322, 5, 88, 45, 2, 318, 322, 5,
	94, 48, 2, 319, 322, 5, 12, 7, 2, 320, 322, 5, 16, 9, 2, 321, 316, 3, 2,
	2, 2, 321, 317, 3, 2, 2, 2, 321, 318, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2,
	321, 320, 3, 2, 2, 2, 322, 49, 3, 2, 2, 2, 323, 333, 5, 44, 23, 2, 324,
	333, 5, 54, 28, 2, 325, 333, 7, 21, 2, 2, 326, 333, 5, 60, 31, 2, 327,
	333, 7, 28, 2, 2, 328, 333, 5, 62, 32, 2, 329, 333, 5, 64, 33, 2, 330,
	333, 5, 66, 34, 2, 331, 333, 5, 68, 35, 2, 332, 323, 3, 2, 2, 2, 332, 324,
	3, 2, 2, 2, 332, 325, 3, 2, 2, 2, 332, 326, 3, 2, 2, 2, 332, 327, 3, 2,
	2, 2, 332, 328, 3, 2, 2, 2, 332, 329, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2,
	332, 331, 3, 2, 2, 2, 333, 51, 3, 2, 2, 2, 334, 335, 9, 9, 2, 2, 335, 53,
	3, 2, 2, 2, 336, 337, 7, 45, 2, 2, 337, 338, 5, 58, 30, 2, 338, 55, 3,
	2, 2, 2, 339, 340, 5, 50, 26, 2, 340, 57, 3, 2, 2, 2, 341, 342, 8, 30,
	1, 2, 342, 343, 5, 56, 29, 2, 343, 349, 3, 2, 2, 2, 344, 345, 12, 3, 2,
	2, 345, 346, 7, 70, 2, 2, 346, 348, 5, 56, 29, 2, 347, 344, 3, 2, 2, 2,
	348, 351, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350,
	59, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 352, 353, 7, 27, 2, 2, 353, 354,
	5, 40, 21, 2, 354, 355, 7, 10, 2, 2, 355, 61, 3, 2, 2, 2, 356, 357, 7,
	29, 2, 2, 357, 358, 5, 44, 23, 2, 358, 63, 3, 2, 2, 2, 359, 362, 7, 30,
	2, 2, 360, 361, 7, 107, 2, 2, 361, 363, 7, 70, 2, 2, 362, 360, 3, 2, 2,
	2, 362, 363, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 365, 7, 107, 2, 2,
	365, 65, 3, 2, 2, 2, 366, 369, 7, 31, 2, 2, 367, 368, 7, 107, 2, 2, 368,
	370, 7, 70, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 371,
	3, 2, 2, 2, 371, 372, 7, 107, 2, 2, 372, 67, 3, 2, 2, 2, 373, 374, 7, 33,
	2, 2, 374, 375, 5, 44, 23, 2, 375, 69, 3, 2, 2, 2, 376, 377, 7, 13, 2,
	2, 377, 378, 7, 107, 2, 2, 378, 379, 7, 65, 2, 2, 379, 380, 5, 76, 39,
	2, 380, 71, 3, 2, 2, 2, 381, 382, 9, 11, 2, 2, 382, 73, 3, 2, 2, 2, 383,
	384, 7, 107, 2, 2, 384, 75, 3, 2, 2, 2, 385, 393, 5, 72, 37, 2, 386, 393,
	5, 78, 40, 2, 387, 393, 5, 82, 42, 2, 388, 393, 5, 84, 43, 2, 389, 393,
	5, 94, 48, 2, 390, 393, 5, 16, 9, 2, 391, 393, 5, 74, 38, 2, 392, 385,
	3, 2, 2, 2, 392, 386, 3, 2, 2, 2, 392, 387, 3, 2, 2, 2, 392, 388, 3, 2,
	2, 2, 392, 389, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 392, 391, 3, 2, 2, 2,
	393, 77, 3, 2, 2, 2, 394, 395, 7, 8, 2, 2, 395, 396, 7, 69, 2, 2, 396,
	397, 7, 8, 2, 2, 397, 79, 3, 2, 2, 2, 398, 399, 8, 41, 1, 2, 399, 400,
	5, 78, 40, 2, 400, 406, 3, 2, 2, 2, 401, 402, 12, 3, 2, 2, 402, 403, 7,
	70, 2, 2, 403, 405, 5, 78, 40, 2, 404, 401, 3, 2, 2, 2, 405, 408, 3, 2,
	2, 2, 406, 404, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 81, 3, 2, 2, 2,
	408, 406, 3, 2, 2, 2, 409, 413, 7, 49, 2, 2, 410, 411, 7, 66, 2, 2, 411,
	412, 7, 8, 2, 2, 412, 414, 7, 67, 2, 2, 413, 410, 3, 2, 2, 2, 413, 414,
	3, 2, 2, 2, 414, 83, 3, 2, 2, 2, 415, 417, 7, 52, 2, 2, 416, 418, 5, 86,
	44, 2, 417, 416, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 417, 3, 2, 2, 2,
	419, 420, 3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 422, 7, 10, 2, 2, 422,
	423, 7, 52, 2, 2, 423, 85, 3, 2, 2, 2, 424, 425, 5, 96, 49, 2, 425, 426,
	7, 65, 2, 2, 426, 427, 5, 76, 39, 2, 427, 87, 3, 2, 2, 2, 428, 429, 7,
	14, 2, 2, 429, 430, 5, 90, 46, 2, 430, 431, 7, 65, 2, 2, 431, 434, 5, 76,
	39, 2, 432, 433, 7, 73, 2, 2, 433, 435, 5, 44, 23, 2, 434, 432, 3, 2, 2,
	2, 434, 435, 3, 2, 2, 2, 435, 442, 3, 2, 2, 2, 436, 437, 7, 14, 2, 2, 437,
	438, 5, 90, 46, 2, 438, 439, 7, 73, 2, 2, 439, 440, 5, 44, 23, 2, 440,
	442, 3, 2, 2, 2, 441, 428, 3, 2, 2, 2, 441, 436, 3, 2, 2, 2, 442, 89, 3,
	2, 2, 2, 443, 448, 5, 92, 47, 2, 444, 445, 7, 70, 2, 2, 445, 447, 5, 92,
	47, 2, 446, 444, 3, 2, 2, 2, 447, 450, 3, 2, 2, 2, 448, 446, 3, 2, 2, 2,
	448, 449, 3, 2, 2, 2, 449, 91, 3, 2, 2, 2, 450, 448, 3, 2, 2, 2, 451, 452,
	7, 107, 2, 2, 452, 93, 3, 2, 2, 2, 453, 454, 7, 50, 2, 2, 454, 455, 5,
	78, 40, 2, 455, 456, 7, 11, 2, 2, 456, 457, 5, 76, 39, 2, 457, 463, 3,
	2, 2, 2, 458, 459, 7, 50, 2, 2, 459, 460, 5, 78, 40, 2, 460, 461, 7, 70,
	2, 2, 461, 463, 3, 2, 2, 2, 462, 453, 3, 2, 2, 2, 462, 458, 3, 2, 2, 2,
	463, 95, 3, 2, 2, 2, 464, 465, 8, 49, 1, 2, 465, 466, 7, 107, 2, 2, 466,
	472, 3, 2, 2, 2, 467, 468, 12, 3, 2, 2, 468, 469, 7, 70, 2, 2, 469, 471,
	7, 107, 2, 2, 470, 467, 3, 2, 2, 2, 471, 474, 3, 2, 2, 2, 472, 470, 3,
	2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 97, 3, 2, 2, 2, 474, 472, 3, 2, 2,
	2, 475, 476, 9, 12, 2, 2, 476, 99, 3, 2, 2, 2, 477, 478, 9, 13, 2, 2, 478,
	101, 3, 2, 2, 2, 46, 105, 109, 115, 118, 124, 127, 130, 136, 140, 154,
	163, 172, 175, 187, 197, 201, 206, 212, 218, 223, 230, 240, 246, 250, 258,
	268, 300, 303, 305, 313, 321, 332, 349, 362, 369, 392, 406, 413, 419, 434,
	441, 448, 462, 472,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'unqualified'", "'pervasive'", "'opaque'", "", "", "", "", "'end'",
	"'of'", "'to'", "'type'", "'var'", "", "", "'class'", "'process'", "'for'",
	"'loop'", "'exit'", "'if'", "'else'", "'elsif'", "'case'", "'assert'",
	"'begin'", "'return'", "'result'", "'new'", "'free'", "'tag'", "'fork'",
	"'signal'", "'wait'", "'pause'", "'quit'", "'unchecked'", "'checked'",
	"'export'", "'import'", "'inherit'", "'implement'", "'by'", "'put'", "'int'",
	"'real'", "'boolean'", "'string'", "'array'", "'set'", "'record'", "'union'",
	"'pointer'", "'nat'", "'intn'", "'natn'", "'realn'", "'char'", "'deferred'",
	"'forward'", "'body'", "'not'", "'^'", "':'", "'('", "')'", "'.'", "'..'",
	"','", "'#'", "'->'", "':='", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'",
	"'rem'", "'**'", "'<'", "'>'", "'='", "'<='", "'>='", "'not='", "'and'",
	"'or'", "'=>'", "'in'", "'shr'", "'shl'", "'xor'", "'+='", "'-='", "'*='",
	"'/='", "'div='", "'mod='", "'and='", "'or='", "'=>='", "'shr='", "'shl='",
	"'xor='",
}
var symbolicNames = []string{
	"", "", "", "", "NOT_IN", "STRING_LITERAL", "INTEGER_LITERAL", "REAL_LITERAL",
	"END", "OF", "TO", "TYPE", "VAR", "PROCEDURE", "FUNCTION", "CLASS", "PROCESS",
	"FOR", "LOOP", "EXIT", "IF", "ELSE", "ELSIF", "CASE", "ASSERT", "BEGIN",
	"RETURN", "RESULT", "NEW", "FREE", "TAG", "FORK", "SIGNAL", "WAIT", "PAUSE",
	"QUIT", "UNCHECKED", "CHECKED", "EXPORT", "IMPORT", "INHERIT", "IMPLEMENT",
	"BY", "PUT", "INT", "REAL", "BOOLEAN", "STRING", "ARRAY", "SET", "RECORD",
	"UNION", "POINTER", "NAT", "INTN", "NATN", "REALN", "CHAR", "DEFERRED",
	"FORWARD", "BODY", "NOT", "CARET", "COLON", "L_PAREN", "R_PAREN", "DOT",
	"RANGE", "COMMA", "CHEAT", "DEREFERENCE", "ASSIGNMENT", "PLUS", "MINUS",
	"MULTIPLY", "DIVIDE", "DIV", "MOD", "REM", "EXP", "LESSTHAN", "GREATERTHAN",
	"EQUAL", "LESSTHANEQUAL", "GREATERTHANEQUAL", "NOTEQUAL", "AND", "OR",
	"IMPLIES", "IN", "SHR", "SHL", "XOR", "PLUSEQUALS", "MINUSEQUALS", "MULTIPLYEQUALS",
	"DIVIDEEQUALS", "DIVEQUALS", "MODEQUALS", "ANDEQUALS", "OREQUALS", "IMPLIESEQUALS",
	"SHREQUALS", "SHLEQUALS", "XOREQUALS", "IDENTIFIER", "WHITESPACE", "BLOCK_COMMENT",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"program", "topLevel", "procHeader", "funcHeader", "subprogramHeader",
	"subprogramDeclaration", "subprogramPrefix", "classDeclaration", "classBody",
	"exportStatement", "exportListItem", "exportList", "inheritStatement",
	"implementStatement", "idOrFileItem", "howExport", "paramDeclaration",
	"paramDeclarationList", "procBody", "statementOrDeclaration", "primaryExpression",
	"expression", "expressionList", "declaration", "statements", "assignmentOperator",
	"putStatement", "putItem", "putItemList", "beginStatement", "resultStatement",
	"newStatement", "freeStatement", "forkStatement", "typeDeclaration", "basicType",
	"referenceType", "typeSpec", "indexType", "indexTypeList", "stringType",
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
	TuringParserNOT_IN           = 4
	TuringParserSTRING_LITERAL   = 5
	TuringParserINTEGER_LITERAL  = 6
	TuringParserREAL_LITERAL     = 7
	TuringParserEND              = 8
	TuringParserOF               = 9
	TuringParserTO               = 10
	TuringParserTYPE             = 11
	TuringParserVAR              = 12
	TuringParserPROCEDURE        = 13
	TuringParserFUNCTION         = 14
	TuringParserCLASS            = 15
	TuringParserPROCESS          = 16
	TuringParserFOR              = 17
	TuringParserLOOP             = 18
	TuringParserEXIT             = 19
	TuringParserIF               = 20
	TuringParserELSE             = 21
	TuringParserELSIF            = 22
	TuringParserCASE             = 23
	TuringParserASSERT           = 24
	TuringParserBEGIN            = 25
	TuringParserRETURN           = 26
	TuringParserRESULT           = 27
	TuringParserNEW              = 28
	TuringParserFREE             = 29
	TuringParserTAG              = 30
	TuringParserFORK             = 31
	TuringParserSIGNAL           = 32
	TuringParserWAIT             = 33
	TuringParserPAUSE            = 34
	TuringParserQUIT             = 35
	TuringParserUNCHECKED        = 36
	TuringParserCHECKED          = 37
	TuringParserEXPORT           = 38
	TuringParserIMPORT           = 39
	TuringParserINHERIT          = 40
	TuringParserIMPLEMENT        = 41
	TuringParserBY               = 42
	TuringParserPUT              = 43
	TuringParserINT              = 44
	TuringParserREAL             = 45
	TuringParserBOOLEAN          = 46
	TuringParserSTRING           = 47
	TuringParserARRAY            = 48
	TuringParserSET              = 49
	TuringParserRECORD           = 50
	TuringParserUNION            = 51
	TuringParserPOINTER          = 52
	TuringParserNAT              = 53
	TuringParserINTN             = 54
	TuringParserNATN             = 55
	TuringParserREALN            = 56
	TuringParserCHAR             = 57
	TuringParserDEFERRED         = 58
	TuringParserFORWARD          = 59
	TuringParserBODY             = 60
	TuringParserNOT              = 61
	TuringParserCARET            = 62
	TuringParserCOLON            = 63
	TuringParserL_PAREN          = 64
	TuringParserR_PAREN          = 65
	TuringParserDOT              = 66
	TuringParserRANGE            = 67
	TuringParserCOMMA            = 68
	TuringParserCHEAT            = 69
	TuringParserDEREFERENCE      = 70
	TuringParserASSIGNMENT       = 71
	TuringParserPLUS             = 72
	TuringParserMINUS            = 73
	TuringParserMULTIPLY         = 74
	TuringParserDIVIDE           = 75
	TuringParserDIV              = 76
	TuringParserMOD              = 77
	TuringParserREM              = 78
	TuringParserEXP              = 79
	TuringParserLESSTHAN         = 80
	TuringParserGREATERTHAN      = 81
	TuringParserEQUAL            = 82
	TuringParserLESSTHANEQUAL    = 83
	TuringParserGREATERTHANEQUAL = 84
	TuringParserNOTEQUAL         = 85
	TuringParserAND              = 86
	TuringParserOR               = 87
	TuringParserIMPLIES          = 88
	TuringParserIN               = 89
	TuringParserSHR              = 90
	TuringParserSHL              = 91
	TuringParserXOR              = 92
	TuringParserPLUSEQUALS       = 93
	TuringParserMINUSEQUALS      = 94
	TuringParserMULTIPLYEQUALS   = 95
	TuringParserDIVIDEEQUALS     = 96
	TuringParserDIVEQUALS        = 97
	TuringParserMODEQUALS        = 98
	TuringParserANDEQUALS        = 99
	TuringParserOREQUALS         = 100
	TuringParserIMPLIESEQUALS    = 101
	TuringParserSHREQUALS        = 102
	TuringParserSHLEQUALS        = 103
	TuringParserXOREQUALS        = 104
	TuringParserIDENTIFIER       = 105
	TuringParserWHITESPACE       = 106
	TuringParserBLOCK_COMMENT    = 107
	TuringParserLINE_COMMENT     = 108
)

// TuringParser rules.
const (
	TuringParserRULE_program                = 0
	TuringParserRULE_topLevel               = 1
	TuringParserRULE_procHeader             = 2
	TuringParserRULE_funcHeader             = 3
	TuringParserRULE_subprogramHeader       = 4
	TuringParserRULE_subprogramDeclaration  = 5
	TuringParserRULE_subprogramPrefix       = 6
	TuringParserRULE_classDeclaration       = 7
	TuringParserRULE_classBody              = 8
	TuringParserRULE_exportStatement        = 9
	TuringParserRULE_exportListItem         = 10
	TuringParserRULE_exportList             = 11
	TuringParserRULE_inheritStatement       = 12
	TuringParserRULE_implementStatement     = 13
	TuringParserRULE_idOrFileItem           = 14
	TuringParserRULE_howExport              = 15
	TuringParserRULE_paramDeclaration       = 16
	TuringParserRULE_paramDeclarationList   = 17
	TuringParserRULE_procBody               = 18
	TuringParserRULE_statementOrDeclaration = 19
	TuringParserRULE_primaryExpression      = 20
	TuringParserRULE_expression             = 21
	TuringParserRULE_expressionList         = 22
	TuringParserRULE_declaration            = 23
	TuringParserRULE_statements             = 24
	TuringParserRULE_assignmentOperator     = 25
	TuringParserRULE_putStatement           = 26
	TuringParserRULE_putItem                = 27
	TuringParserRULE_putItemList            = 28
	TuringParserRULE_beginStatement         = 29
	TuringParserRULE_resultStatement        = 30
	TuringParserRULE_newStatement           = 31
	TuringParserRULE_freeStatement          = 32
	TuringParserRULE_forkStatement          = 33
	TuringParserRULE_typeDeclaration        = 34
	TuringParserRULE_basicType              = 35
	TuringParserRULE_referenceType          = 36
	TuringParserRULE_typeSpec               = 37
	TuringParserRULE_indexType              = 38
	TuringParserRULE_indexTypeList          = 39
	TuringParserRULE_stringType             = 40
	TuringParserRULE_recordType             = 41
	TuringParserRULE_recordField            = 42
	TuringParserRULE_variableDeclaration    = 43
	TuringParserRULE_variableIdentifierList = 44
	TuringParserRULE_variableIdentifier     = 45
	TuringParserRULE_arrayDeclaration       = 46
	TuringParserRULE_identifierList         = 47
	TuringParserRULE_literal                = 48
	TuringParserRULE_comment                = 49
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

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserREAL_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserCLASS)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(TuringParserPUT-43))|(1<<(TuringParserARRAY-43))|(1<<(TuringParserDEFERRED-43))|(1<<(TuringParserFORWARD-43))|(1<<(TuringParserBODY-43))|(1<<(TuringParserNOT-43))|(1<<(TuringParserCARET-43))|(1<<(TuringParserL_PAREN-43))|(1<<(TuringParserCHEAT-43))|(1<<(TuringParserPLUS-43))|(1<<(TuringParserMINUS-43)))) != 0) || _la == TuringParserIDENTIFIER {
		{
			p.SetState(100)
			p.TopLevel()
		}

		p.SetState(103)
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

func (s *TopLevelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitTopLevel(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.StatementOrDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.ClassDeclaration()
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

func (s *ProcHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitProcHeader(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(109)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TuringParserPROCEDURE || _la == TuringParserPROCESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(110)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(111)
			p.Match(TuringParserL_PAREN)
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(TuringParserDEFERRED-58))|(1<<(TuringParserFORWARD-58))|(1<<(TuringParserBODY-58)))) != 0) || _la == TuringParserIDENTIFIER {
			{
				p.SetState(112)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(115)
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

func (s *FuncHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitFuncHeader(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(118)
		p.Match(TuringParserFUNCTION)
	}
	{
		p.SetState(119)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TuringParserL_PAREN {
		{
			p.SetState(120)
			p.Match(TuringParserL_PAREN)
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserPROCESS))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(TuringParserDEFERRED-58))|(1<<(TuringParserFORWARD-58))|(1<<(TuringParserBODY-58)))) != 0) || _la == TuringParserIDENTIFIER {
			{
				p.SetState(121)
				p.paramDeclarationList(0)
			}

		}
		{
			p.SetState(124)
			p.Match(TuringParserR_PAREN)
		}

	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TuringParserIDENTIFIER {
		{
			p.SetState(127)
			p.Match(TuringParserIDENTIFIER)
		}

	}
	{
		p.SetState(130)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(131)
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

func (s *SubprogramHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitSubprogramHeader(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(TuringParserDEFERRED-58))|(1<<(TuringParserFORWARD-58))|(1<<(TuringParserBODY-58)))) != 0 {
		{
			p.SetState(133)
			p.SubprogramPrefix()
		}

	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserFUNCTION:
		{
			p.SetState(136)
			p.FuncHeader()
		}

	case TuringParserPROCEDURE, TuringParserPROCESS:
		{
			p.SetState(137)
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

func (s *SubprogramDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitSubprogramDeclaration(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(140)
		p.SubprogramHeader()
	}
	{
		p.SetState(141)
		p.ProcBody()
	}
	{
		p.SetState(142)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(143)
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

func (s *SubprogramPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitSubprogramPrefix(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(145)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(TuringParserDEFERRED-58))|(1<<(TuringParserFORWARD-58))|(1<<(TuringParserBODY-58)))) != 0) {
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

func (s *ClassDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitClassDeclaration(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(147)
		p.Match(TuringParserCLASS)
	}
	{
		p.SetState(148)
		p.Match(TuringParserIDENTIFIER)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserREAL_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserCLASS)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(TuringParserEXPORT-38))|(1<<(TuringParserINHERIT-38))|(1<<(TuringParserIMPLEMENT-38))|(1<<(TuringParserPUT-38))|(1<<(TuringParserARRAY-38))|(1<<(TuringParserDEFERRED-38))|(1<<(TuringParserFORWARD-38))|(1<<(TuringParserBODY-38))|(1<<(TuringParserNOT-38))|(1<<(TuringParserCARET-38))|(1<<(TuringParserL_PAREN-38))|(1<<(TuringParserCHEAT-38)))) != 0) || _la == TuringParserPLUS || _la == TuringParserMINUS || _la == TuringParserIDENTIFIER {
		{
			p.SetState(149)
			p.ClassBody()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(154)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(155)
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

func (s *ClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitClassBody(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserEXPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.ExportStatement()
		}

	case TuringParserINHERIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.InheritStatement()
		}

	case TuringParserIMPLEMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.ImplementStatement()
		}

	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserREAL_LITERAL, TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserCLASS, TuringParserPROCESS, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
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

func (s *ExportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitExportStatement(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.Match(TuringParserEXPORT)
		}
		{
			p.SetState(164)
			p.exportList(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(TuringParserEXPORT)
		}
		{
			p.SetState(166)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(167)
			p.exportList(0)
		}
		{
			p.SetState(168)
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

func (s *ExportListItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitExportListItem(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserT__0)|(1<<TuringParserT__1)|(1<<TuringParserT__2)|(1<<TuringParserVAR))) != 0 {
		{
			p.SetState(172)
			p.HowExport()
		}

	}
	{
		p.SetState(175)
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

func (s *ExportListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitExportList(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(178)
		p.ExportListItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(185)
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
			p.SetState(180)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(181)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(182)
				p.ExportListItem()
			}

		}
		p.SetState(187)
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

func (s *InheritStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitInheritStatement(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(TuringParserINHERIT)
		}
		{
			p.SetState(189)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(TuringParserINHERIT)
		}
		{
			p.SetState(191)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(192)
			p.IdOrFileItem()
		}
		{
			p.SetState(193)
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

func (s *ImplementStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitImplementStatement(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(TuringParserIMPLEMENT)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserBY {
			{
				p.SetState(198)
				p.Match(TuringParserBY)
			}

		}
		{
			p.SetState(201)
			p.IdOrFileItem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(TuringParserIMPLEMENT)
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserBY {
			{
				p.SetState(203)
				p.Match(TuringParserBY)
			}

		}
		{
			p.SetState(206)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(207)
			p.IdOrFileItem()
		}
		{
			p.SetState(208)
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

func (s *IdOrFileItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitIdOrFileItem(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(TuringParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(214)
			p.Match(TuringParserIN)
		}
		{
			p.SetState(215)
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

func (s *HowExportContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitHowExport(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(218)
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

func (s *ParamDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitParamDeclaration(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserVAR, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserVAR {
			{
				p.SetState(220)
				p.Match(TuringParserVAR)
			}

		}
		{
			p.SetState(223)
			p.identifierList(0)
		}
		{
			p.SetState(224)
			p.Match(TuringParserCOLON)
		}
		{
			p.SetState(225)
			p.TypeSpec()
		}

	case TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
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

func (s *ParamDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitParamDeclarationList(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(231)
		p.ParamDeclaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(238)
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
			p.SetState(233)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(234)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(235)
				p.ParamDeclaration()
			}

		}
		p.SetState(240)
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

func (s *ProcBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitProcBody(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserREAL_LITERAL)|(1<<TuringParserTYPE)|(1<<TuringParserVAR)|(1<<TuringParserPROCEDURE)|(1<<TuringParserFUNCTION)|(1<<TuringParserCLASS)|(1<<TuringParserPROCESS)|(1<<TuringParserEXIT)|(1<<TuringParserBEGIN)|(1<<TuringParserRETURN)|(1<<TuringParserRESULT)|(1<<TuringParserNEW)|(1<<TuringParserFREE)|(1<<TuringParserFORK))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(TuringParserPUT-43))|(1<<(TuringParserARRAY-43))|(1<<(TuringParserDEFERRED-43))|(1<<(TuringParserFORWARD-43))|(1<<(TuringParserBODY-43))|(1<<(TuringParserNOT-43))|(1<<(TuringParserCARET-43))|(1<<(TuringParserL_PAREN-43))|(1<<(TuringParserCHEAT-43))|(1<<(TuringParserPLUS-43))|(1<<(TuringParserMINUS-43)))) != 0) || _la == TuringParserIDENTIFIER {
		{
			p.SetState(241)
			p.StatementOrDeclaration()
		}

		p.SetState(244)
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

func (s *StatementOrDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitStatementOrDeclaration(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserREAL_LITERAL, TuringParserEXIT, TuringParserBEGIN, TuringParserRETURN, TuringParserRESULT, TuringParserNEW, TuringParserFREE, TuringParserFORK, TuringParserPUT, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.Statements()
		}

	case TuringParserTYPE, TuringParserVAR, TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserCLASS, TuringParserPROCESS, TuringParserARRAY, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
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

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Literal()
		}

	case TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(TuringParserIDENTIFIER)
		}

	case TuringParserL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(252)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(253)
			p.expression(0)
		}
		{
			p.SetState(254)
			p.Match(TuringParserR_PAREN)
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
	return s.GetToken(TuringParserCARET, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TuringParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TuringParserMINUS, 0)
}

func (s *ExpressionContext) CHEAT() antlr.TerminalNode {
	return s.GetToken(TuringParserCHEAT, 0)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TuringParserNOT, 0)
}

func (s *ExpressionContext) EXP() antlr.TerminalNode {
	return s.GetToken(TuringParserEXP, 0)
}

func (s *ExpressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(TuringParserMULTIPLY, 0)
}

func (s *ExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVIDE, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(TuringParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(TuringParserMOD, 0)
}

func (s *ExpressionContext) REM() antlr.TerminalNode {
	return s.GetToken(TuringParserREM, 0)
}

func (s *ExpressionContext) SHR() antlr.TerminalNode {
	return s.GetToken(TuringParserSHR, 0)
}

func (s *ExpressionContext) SHL() antlr.TerminalNode {
	return s.GetToken(TuringParserSHL, 0)
}

func (s *ExpressionContext) XOR() antlr.TerminalNode {
	return s.GetToken(TuringParserXOR, 0)
}

func (s *ExpressionContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(TuringParserLESSTHAN, 0)
}

func (s *ExpressionContext) GREATERTHAN() antlr.TerminalNode {
	return s.GetToken(TuringParserGREATERTHAN, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserEQUAL, 0)
}

func (s *ExpressionContext) LESSTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserLESSTHANEQUAL, 0)
}

func (s *ExpressionContext) GREATERTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserGREATERTHANEQUAL, 0)
}

func (s *ExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(TuringParserNOTEQUAL, 0)
}

func (s *ExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(TuringParserIN, 0)
}

func (s *ExpressionContext) NOT_IN() antlr.TerminalNode {
	return s.GetToken(TuringParserNOT_IN, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(TuringParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(TuringParserOR, 0)
}

func (s *ExpressionContext) IMPLIES() antlr.TerminalNode {
	return s.GetToken(TuringParserIMPLIES, 0)
}

func (s *ExpressionContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(TuringParserASSIGNMENT, 0)
}

func (s *ExpressionContext) PLUSEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserPLUSEQUALS, 0)
}

func (s *ExpressionContext) MINUSEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserMINUSEQUALS, 0)
}

func (s *ExpressionContext) MULTIPLYEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserMULTIPLYEQUALS, 0)
}

func (s *ExpressionContext) DIVIDEEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVIDEEQUALS, 0)
}

func (s *ExpressionContext) DIVEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVEQUALS, 0)
}

func (s *ExpressionContext) SHLEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserSHLEQUALS, 0)
}

func (s *ExpressionContext) SHREQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserSHREQUALS, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *ExpressionContext) DEREFERENCE() antlr.TerminalNode {
	return s.GetToken(TuringParserDEREFERENCE, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(TuringParserDOT, 0)
}

func (s *ExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserL_PAREN, 0)
}

func (s *ExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TuringParserR_PAREN, 0)
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
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TuringParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, TuringParserRULE_expression, _p)
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
	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserREAL_LITERAL, TuringParserL_PAREN, TuringParserIDENTIFIER:
		{
			p.SetState(259)
			p.PrimaryExpression()
		}

	case TuringParserCARET:
		{
			p.SetState(260)

			var _m = p.Match(TuringParserCARET)

			localctx.(*ExpressionContext).prefix = _m
		}
		{
			p.SetState(261)
			p.expression(13)
		}

	case TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS:
		{
			p.SetState(262)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(TuringParserCHEAT-69))|(1<<(TuringParserPLUS-69))|(1<<(TuringParserMINUS-69)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(263)
			p.expression(9)
		}

	case TuringParserNOT:
		{
			p.SetState(264)

			var _m = p.Match(TuringParserNOT)

			localctx.(*ExpressionContext).prefix = _m
		}
		{
			p.SetState(265)
			p.expression(5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(301)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(269)

					var _m = p.Match(TuringParserEXP)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(270)
					p.expression(12)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(272)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-74)&-(0x1f+1)) == 0 && ((1<<uint((_la-74)))&((1<<(TuringParserMULTIPLY-74))|(1<<(TuringParserDIVIDE-74))|(1<<(TuringParserDIV-74))|(1<<(TuringParserMOD-74))|(1<<(TuringParserREM-74))|(1<<(TuringParserSHR-74))|(1<<(TuringParserSHL-74)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(273)
					p.expression(9)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(275)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(TuringParserPLUS-72))|(1<<(TuringParserMINUS-72))|(1<<(TuringParserXOR-72)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(276)
					p.expression(8)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(278)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TuringParserNOT_IN || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(TuringParserLESSTHAN-80))|(1<<(TuringParserGREATERTHAN-80))|(1<<(TuringParserEQUAL-80))|(1<<(TuringParserLESSTHANEQUAL-80))|(1<<(TuringParserGREATERTHANEQUAL-80))|(1<<(TuringParserNOTEQUAL-80))|(1<<(TuringParserIN-80)))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(279)
					p.expression(7)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(281)

					var _m = p.Match(TuringParserAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(282)
					p.expression(5)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(284)

					var _m = p.Match(TuringParserOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(285)
					p.expression(4)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(287)

					var _m = p.Match(TuringParserIMPLIES)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(288)
					p.expression(3)
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(290)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(TuringParserASSIGNMENT-71))|(1<<(TuringParserPLUSEQUALS-71))|(1<<(TuringParserMINUSEQUALS-71))|(1<<(TuringParserMULTIPLYEQUALS-71))|(1<<(TuringParserDIVIDEEQUALS-71))|(1<<(TuringParserDIVEQUALS-71))|(1<<(TuringParserSHREQUALS-71)))) != 0) || _la == TuringParserSHLEQUALS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(291)
					p.expression(1)
				}

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(293)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TuringParserDOT || _la == TuringParserDEREFERENCE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(294)
					p.Match(TuringParserIDENTIFIER)
				}

			case 10:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_expression)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(296)
					p.Match(TuringParserL_PAREN)
				}
				p.SetState(298)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserREAL_LITERAL))) != 0) || (((_la-61)&-(0x1f+1)) == 0 && ((1<<uint((_la-61)))&((1<<(TuringParserNOT-61))|(1<<(TuringParserCARET-61))|(1<<(TuringParserL_PAREN-61))|(1<<(TuringParserCHEAT-61))|(1<<(TuringParserPLUS-61))|(1<<(TuringParserMINUS-61)))) != 0) || _la == TuringParserIDENTIFIER {
					{
						p.SetState(297)
						p.ExpressionList()
					}

				}
				{
					p.SetState(300)
					p.Match(TuringParserR_PAREN)
				}

			}

		}
		p.SetState(305)
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
	return s.GetTokens(TuringParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, i)
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

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TuringParserRULE_expressionList)
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
		p.SetState(306)
		p.expression(0)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TuringParserCOMMA {
		{
			p.SetState(307)
			p.Match(TuringParserCOMMA)
		}
		{
			p.SetState(308)
			p.expression(0)
		}

		p.SetState(313)
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
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TuringParserRULE_declaration)

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

	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)
			p.TypeDeclaration()
		}

	case TuringParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.VariableDeclaration()
		}

	case TuringParserARRAY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(316)
			p.ArrayDeclaration()
		}

	case TuringParserPROCEDURE, TuringParserFUNCTION, TuringParserPROCESS, TuringParserDEFERRED, TuringParserFORWARD, TuringParserBODY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(317)
			p.SubprogramDeclaration()
		}

	case TuringParserCLASS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(318)
			p.ClassDeclaration()
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

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TuringParserRULE_statements)

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

	p.SetState(330)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserSTRING_LITERAL, TuringParserINTEGER_LITERAL, TuringParserREAL_LITERAL, TuringParserNOT, TuringParserCARET, TuringParserL_PAREN, TuringParserCHEAT, TuringParserPLUS, TuringParserMINUS, TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.expression(0)
		}

	case TuringParserPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.PutStatement()
		}

	case TuringParserEXIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(323)
			p.Match(TuringParserEXIT)
		}

	case TuringParserBEGIN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(324)
			p.BeginStatement()
		}

	case TuringParserRETURN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(325)
			p.Match(TuringParserRETURN)
		}

	case TuringParserRESULT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(326)
			p.ResultStatement()
		}

	case TuringParserNEW:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(327)
			p.NewStatement()
		}

	case TuringParserFREE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(328)
			p.FreeStatement()
		}

	case TuringParserFORK:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(329)
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
	p.RuleIndex = TuringParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(TuringParserASSIGNMENT, 0)
}

func (s *AssignmentOperatorContext) PLUSEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserPLUSEQUALS, 0)
}

func (s *AssignmentOperatorContext) MINUSEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserMINUSEQUALS, 0)
}

func (s *AssignmentOperatorContext) MULTIPLYEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserMULTIPLYEQUALS, 0)
}

func (s *AssignmentOperatorContext) DIVIDEEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVIDEEQUALS, 0)
}

func (s *AssignmentOperatorContext) DIVEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserDIVEQUALS, 0)
}

func (s *AssignmentOperatorContext) SHLEQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserSHLEQUALS, 0)
}

func (s *AssignmentOperatorContext) SHREQUALS() antlr.TerminalNode {
	return s.GetToken(TuringParserSHREQUALS, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitAssignmentOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TuringParserRULE_assignmentOperator)
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
		p.SetState(332)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(TuringParserASSIGNMENT-71))|(1<<(TuringParserPLUSEQUALS-71))|(1<<(TuringParserMINUSEQUALS-71))|(1<<(TuringParserMULTIPLYEQUALS-71))|(1<<(TuringParserDIVIDEEQUALS-71))|(1<<(TuringParserDIVEQUALS-71))|(1<<(TuringParserSHREQUALS-71)))) != 0) || _la == TuringParserSHLEQUALS) {
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

func (s *PutStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitPutStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) PutStatement() (localctx IPutStatementContext) {
	localctx = NewPutStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TuringParserRULE_putStatement)

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
		p.SetState(334)
		p.Match(TuringParserPUT)
	}
	{
		p.SetState(335)
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

func (s *PutItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitPutItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) PutItem() (localctx IPutItemContext) {
	localctx = NewPutItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TuringParserRULE_putItem)

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

func (s *PutItemListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitPutItemList(s)

	default:
		return t.VisitChildren(s)
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
	_startState := 56
	p.EnterRecursionRule(localctx, 56, TuringParserRULE_putItemList, _p)

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
		p.SetState(340)
		p.PutItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPutItemListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_putItemList)
			p.SetState(342)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(343)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(344)
				p.PutItem()
			}

		}
		p.SetState(349)
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

func (s *BeginStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitBeginStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) BeginStatement() (localctx IBeginStatementContext) {
	localctx = NewBeginStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TuringParserRULE_beginStatement)

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
		p.SetState(350)
		p.Match(TuringParserBEGIN)
	}
	{
		p.SetState(351)
		p.StatementOrDeclaration()
	}
	{
		p.SetState(352)
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

func (s *ResultStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitResultStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) ResultStatement() (localctx IResultStatementContext) {
	localctx = NewResultStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TuringParserRULE_resultStatement)

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
		p.SetState(354)
		p.Match(TuringParserRESULT)
	}
	{
		p.SetState(355)
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

func (s *NewStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitNewStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) NewStatement() (localctx INewStatementContext) {
	localctx = NewNewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, TuringParserRULE_newStatement)

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
		p.Match(TuringParserNEW)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(358)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(359)
			p.Match(TuringParserCOMMA)
		}

	}
	{
		p.SetState(362)
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

func (s *FreeStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitFreeStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) FreeStatement() (localctx IFreeStatementContext) {
	localctx = NewFreeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, TuringParserRULE_freeStatement)

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
		p.SetState(364)
		p.Match(TuringParserFREE)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(365)
			p.Match(TuringParserIDENTIFIER)
		}
		{
			p.SetState(366)
			p.Match(TuringParserCOMMA)
		}

	}
	{
		p.SetState(369)
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

func (s *ForkStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitForkStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) ForkStatement() (localctx IForkStatementContext) {
	localctx = NewForkStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, TuringParserRULE_forkStatement)

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
		p.SetState(371)
		p.Match(TuringParserFORK)
	}
	{
		p.SetState(372)
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

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, TuringParserRULE_typeDeclaration)

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
		p.Match(TuringParserTYPE)
	}
	{
		p.SetState(375)
		p.Match(TuringParserIDENTIFIER)
	}
	{
		p.SetState(376)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(377)
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
	p.RuleIndex = TuringParserRULE_basicType
	return p
}

func (*BasicTypeContext) IsBasicTypeContext() {}

func NewBasicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicTypeContext {
	var p = new(BasicTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_basicType

	return p
}

func (s *BasicTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(TuringParserINT, 0)
}

func (s *BasicTypeContext) REAL() antlr.TerminalNode {
	return s.GetToken(TuringParserREAL, 0)
}

func (s *BasicTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(TuringParserBOOLEAN, 0)
}

func (s *BasicTypeContext) NAT() antlr.TerminalNode {
	return s.GetToken(TuringParserNAT, 0)
}

func (s *BasicTypeContext) INTN() antlr.TerminalNode {
	return s.GetToken(TuringParserINTN, 0)
}

func (s *BasicTypeContext) NATN() antlr.TerminalNode {
	return s.GetToken(TuringParserNATN, 0)
}

func (s *BasicTypeContext) REALN() antlr.TerminalNode {
	return s.GetToken(TuringParserREALN, 0)
}

func (s *BasicTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(TuringParserCHAR, 0)
}

func (s *BasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterBasicType(s)
	}
}

func (s *BasicTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitBasicType(s)
	}
}

func (s *BasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) BasicType() (localctx IBasicTypeContext) {
	localctx = NewBasicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, TuringParserRULE_basicType)
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
		p.SetState(379)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(TuringParserINT-44))|(1<<(TuringParserREAL-44))|(1<<(TuringParserBOOLEAN-44))|(1<<(TuringParserNAT-44))|(1<<(TuringParserINTN-44))|(1<<(TuringParserNATN-44))|(1<<(TuringParserREALN-44))|(1<<(TuringParserCHAR-44)))) != 0) {
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
	p.RuleIndex = TuringParserRULE_referenceType
	return p
}

func (*ReferenceTypeContext) IsReferenceTypeContext() {}

func NewReferenceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceTypeContext {
	var p = new(ReferenceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_referenceType

	return p
}

func (s *ReferenceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TuringParserIDENTIFIER, 0)
}

func (s *ReferenceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterReferenceType(s)
	}
}

func (s *ReferenceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitReferenceType(s)
	}
}

func (s *ReferenceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitReferenceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) ReferenceType() (localctx IReferenceTypeContext) {
	localctx = NewReferenceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, TuringParserRULE_referenceType)

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
		p.SetState(381)
		p.Match(TuringParserIDENTIFIER)
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
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, TuringParserRULE_typeSpec)

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

	p.SetState(390)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TuringParserINT, TuringParserREAL, TuringParserBOOLEAN, TuringParserNAT, TuringParserINTN, TuringParserNATN, TuringParserREALN, TuringParserCHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.BasicType()
		}

	case TuringParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.IndexType()
		}

	case TuringParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(385)
			p.StringType()
		}

	case TuringParserRECORD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(386)
			p.RecordType()
		}

	case TuringParserARRAY:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(387)
			p.ArrayDeclaration()
		}

	case TuringParserCLASS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(388)
			p.ClassDeclaration()
		}

	case TuringParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(389)
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

func (s *IndexTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitIndexType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) IndexType() (localctx IIndexTypeContext) {
	localctx = NewIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, TuringParserRULE_indexType)

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
		p.SetState(392)
		p.Match(TuringParserINTEGER_LITERAL)
	}
	{
		p.SetState(393)
		p.Match(TuringParserRANGE)
	}
	{
		p.SetState(394)
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

func (s *IndexTypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitIndexTypeList(s)

	default:
		return t.VisitChildren(s)
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
	_startState := 78
	p.EnterRecursionRule(localctx, 78, TuringParserRULE_indexTypeList, _p)

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
		p.IndexType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIndexTypeListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_indexTypeList)
			p.SetState(399)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(400)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(401)
				p.IndexType()
			}

		}
		p.SetState(406)
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

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) StringType() (localctx IStringTypeContext) {
	localctx = NewStringTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, TuringParserRULE_stringType)

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
		p.SetState(407)
		p.Match(TuringParserSTRING)
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(408)
			p.Match(TuringParserL_PAREN)
		}
		{
			p.SetState(409)
			p.Match(TuringParserINTEGER_LITERAL)
		}
		{
			p.SetState(410)
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

func (s *RecordTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitRecordType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) RecordType() (localctx IRecordTypeContext) {
	localctx = NewRecordTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, TuringParserRULE_recordType)
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
		p.SetState(413)
		p.Match(TuringParserRECORD)
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TuringParserIDENTIFIER {
		{
			p.SetState(414)
			p.RecordField()
		}

		p.SetState(417)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(419)
		p.Match(TuringParserEND)
	}
	{
		p.SetState(420)
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

func (s *RecordFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitRecordField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) RecordField() (localctx IRecordFieldContext) {
	localctx = NewRecordFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, TuringParserRULE_recordField)

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
		p.SetState(422)
		p.identifierList(0)
	}
	{
		p.SetState(423)
		p.Match(TuringParserCOLON)
	}
	{
		p.SetState(424)
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

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, TuringParserRULE_variableDeclaration)
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

	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(426)
			p.Match(TuringParserVAR)
		}
		{
			p.SetState(427)
			p.VariableIdentifierList()
		}
		{
			p.SetState(428)
			p.Match(TuringParserCOLON)
		}
		{
			p.SetState(429)
			p.TypeSpec()
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TuringParserASSIGNMENT {
			{
				p.SetState(430)
				p.Match(TuringParserASSIGNMENT)
			}
			{
				p.SetState(431)
				p.expression(0)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)
			p.Match(TuringParserVAR)
		}
		{
			p.SetState(435)
			p.VariableIdentifierList()
		}
		{
			p.SetState(436)
			p.Match(TuringParserASSIGNMENT)
		}
		{
			p.SetState(437)
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
	return s.GetTokens(TuringParserCOMMA)
}

func (s *VariableIdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TuringParserCOMMA, i)
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

func (s *VariableIdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitVariableIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) VariableIdentifierList() (localctx IVariableIdentifierListContext) {
	localctx = NewVariableIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, TuringParserRULE_variableIdentifierList)
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
		p.SetState(441)
		p.VariableIdentifier()
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TuringParserCOMMA {
		{
			p.SetState(442)
			p.Match(TuringParserCOMMA)
		}
		{
			p.SetState(443)
			p.VariableIdentifier()
		}

		p.SetState(448)
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

func (s *VariableIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitVariableIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) VariableIdentifier() (localctx IVariableIdentifierContext) {
	localctx = NewVariableIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, TuringParserRULE_variableIdentifier)

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

func (s *ArrayDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitArrayDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) ArrayDeclaration() (localctx IArrayDeclarationContext) {
	localctx = NewArrayDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, TuringParserRULE_arrayDeclaration)

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

	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(451)
			p.Match(TuringParserARRAY)
		}
		{
			p.SetState(452)
			p.IndexType()
		}
		{
			p.SetState(453)
			p.Match(TuringParserOF)
		}
		{
			p.SetState(454)
			p.TypeSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(456)
			p.Match(TuringParserARRAY)
		}
		{
			p.SetState(457)
			p.IndexType()
		}
		{
			p.SetState(458)
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

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
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
	_startState := 94
	p.EnterRecursionRule(localctx, 94, TuringParserRULE_identifierList, _p)

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
		p.SetState(463)
		p.Match(TuringParserIDENTIFIER)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIdentifierListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TuringParserRULE_identifierList)
			p.SetState(465)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(466)
				p.Match(TuringParserCOMMA)
			}
			{
				p.SetState(467)
				p.Match(TuringParserIDENTIFIER)
			}

		}
		p.SetState(472)
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

func (s *LiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(TuringParserREAL_LITERAL, 0)
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

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, TuringParserRULE_literal)
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
		p.SetState(473)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TuringParserSTRING_LITERAL)|(1<<TuringParserINTEGER_LITERAL)|(1<<TuringParserREAL_LITERAL))) != 0) {
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

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TuringVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TuringParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, TuringParserRULE_comment)
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
		p.SetState(475)
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

func (p *TuringParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

func (p *TuringParser) PutItemList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) IndexTypeList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TuringParser) IdentifierList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
