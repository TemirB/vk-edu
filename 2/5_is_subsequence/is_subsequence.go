package issubsequence

import "container/list"

/*
	Является ли одна строка исходной для другой

	В исходную строку добавили некоторое количество символов.
	Необходимо выявить, была ли строка a исходной для строки b.

	Не написал ноды на дженериках, теперь жалею,
	пусть тут пока будет из стандратного пакета
*/

func IsSubsequence(a, b string) bool {
	if a == "" {
		return true
	}
	if b == "" {
		return false
	}

	l := list.New()
	for _, r := range a {
		l.PushBack(r)
	}

	for _, r := range b {
		front := l.Front()
		if front == nil {
			// очередь пустая = все символы найдены
			return true
		}

		if front.Value.(rune) == r {
			l.Remove(front)
		}
	}

	return l.Len() == 0
}
