import sys
import re


class Stack:
	def __init__(self):
		self.items = []

	def isEmpty(self):
		return self.items == []

	def push(self, item):
		self.items.append(item)

	def pop(self):
		return self.items.pop()

	def peek(self):
		return self.items[len(self.items) - 1]

	def size(self):
		return len(self.items)


class Queue:
	def __init__(self):
		self.items = []

	def isEmpty(self):
		return self.items == []

	def push(self, item):
		self.items.insert(0, item)

	def pop(self):
		return self.items.pop()

	def size(self):
		return len(self.items)


def traverse(first_vertex, edges, isBreadth):
	used = set()

	container = Queue() if isBreadth else Stack()
	container.push(first_vertex)

	while container:
		if container.isEmpty():
			break

		v = container.pop()
		if v in used: continue

		used.add(v)
		print(v)
		neighbors = edges.get(v)
		if neighbors:
			for neighbor in (neighbors if isBreadth else neighbors[::-1]):
				container.push(neighbor)


def addToContainer(container, key, val):
	if container.get(key):
		container[key].add(val)
	else:
		container[key] = {val}


def main():
	p = re.compile('(?P<init>^[ud] \S+ [bd]$)|(?P<edge>^\S+ \S+$)|(?P<empty>^$)')
	edges = None

	for line in sys.stdin:
		line = line[:len(line) - 1]
		m = p.fullmatch(line)

		if not m:
			print('error')
			continue

		if m.group('empty'):
			continue

		words = line.split(' ')

		if m.group('init'):
			if edges is not None:
				print('error')
			isOriented = words[0] == 'd'
			v = words[1]
			isBreadth = words[2] == 'b'
			edges = {}

		elif m.group('edge'):
			if edges is None:
				print('error')

			addToContainer(edges, words[0], words[1])
			if not isOriented:
				addToContainer(edges, words[1], words[0])

	if edges is None:
		print('error')
		return

	for key, vertices in edges.items():
		edges[key] = sorted(vertices)

	traverse(v, edges, isBreadth)


if __name__ == '__main__':
	main()
