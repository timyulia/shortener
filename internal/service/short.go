package service

import (
	"crypto/sha256"
	"strings"
)

const (
	forbidden = 63
	decision  = 45
	alphabet  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

// generateShortURL
/*
Алгоритм на основе хеш-функции и base-63.
Генерируется хеш-сумма длиной 224 по входной строке (изначально по длинной ссылке). Затем от кадого из
десяти первых байтов берутся 6 младших бит, которые вместе образуют значения от нуля до 64.
Мощность алфавита, из которого состоят ссылки, - 63. Значит, этим алфавитом можно закодировать 63 числовых значения.
Таким образом, есть одно лишнее значение, образуемое шестью битами - 63 (ему нет соответствующего символа в алфавите).
Если запрещенное значение 63 попадается, оно последовательно ксорится с байтами (с 10 до 27), пока не изменится на другое (достаточно
одного раза, кроме случаев когда попались 6 нулевых бит, но для обработки исключительных случаев делается в цикле).
*/
func generateShortURL(long string) string {
	hash := sha256.Sum224([]byte(long))
	res := strings.Builder{}
	j := 0
	for i := 0; i < 10; i++ {
		curr := hash[i] & ((1 << 6) - 1)
		jOld := j
		for curr == forbidden {
			curr ^= hash[j+10] & ((1 << 6) - 1)
			j = (j + 1) % 18
			if j == jOld && curr == forbidden { //если младшие 6 бит всех байт с 10 до 27 равны нулям (вероятность чего почти нулевая), необходимо выйти из цикла добавлением константы
				curr ^= decision
			}
		}
		symbol := alphabet[curr]
		res.WriteByte(symbol)
	}
	return res.String()
}
