#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <list>
#include <stdexcept>
#include <algorithm>
#include <ctime>

#define methods
#define parametres

namespace asd {
    template<class T>
    class Stack {
    public methods:

        explicit Stack(const int size = 10)
                : _size(size), _top(-1), _buffer(new T[size]) {}

        ~Stack() {
            delete[] _buffer;
        }

        void push(const T value) {
            _buffer[++_top] = value;
        }

        void pop() {
            --_top;
        }

        T &top() const {
            return _buffer[_top];
        }

        bool isEmpty() const {
            return _top == -1;
        }

        bool isFull() const {
            return _size == _top + 1;
        }

        int size() const {
            return _top;
        }

        void print(std::ostream &output) {
            for (int i = _top; i >= 0; --i) {
                output << _buffer[i] << std::endl;
            }
        }

    private parametres:
        T *_buffer;
        int _size;
        int _top;
    };

    template<class T>
    class Queue {
    public methods:

        explicit Queue(const int capacity = 10)
                : _capacity(capacity), _head(-1), _tail(-1), _buffer(new T[capacity]) {}

        ~Queue() {
            delete[] _buffer;
        }

        void push(const T &value) {
            if (isEmpty()) {
                _head = _tail = 0;
            } else {
                _tail = (_tail + 1) % capacity();
            }

            _buffer[_tail] = value;
        }

        void pop() {
            if (_head == _tail) {
                _head = _tail = -1;
            } else {
                _head = (_head + 1) % capacity();
            }
        }

        T top() const {
            return _buffer[_head];
        }

        bool isEmpty() const {
            return _head == -1 && _tail == -1;
        }

        bool isFull() const {
            return (_tail + 1) % capacity() == _head;
        }

        int capacity() const {
            return _capacity;
        }

        void print(std::ostream &output) {
            const int length = (_tail - _head + capacity()) % capacity() + 1;

            for (int i = length; i >= 0; --i) {
                output << _buffer[(_head + i) % capacity()] << std::endl;
            }
        }

    private parametres:
        T *_buffer;
        int _capacity;
        int _head;
        int _tail;
    };

    class Graph {
    public methods:

        Graph(const std::size_t vertices)
                : _vertices(vertices), _adj(new std::vector<int>[vertices]) {}

        virtual ~Graph() {
            delete[] _adj;
        }

        virtual void addEdge(int v, int w) {}

        void DFSRecursive(int v) {
            std::vector<bool> visited(_vertices, false);
            DFSUtil(v, visited);
        }

        void DFS(int start) {
            std::vector<bool> visited(_vertices, false);
            Stack<int> stack;
            Stack<int> result;

            stack.push(start);
            visited[start] = true;

            while (!stack.isEmpty()) {
                for (auto it = _adj[start].begin(); it != _adj[start].end(); ++it) {
                    if (!visited[*it]) {
                        visited[*it] = true;
                        stack.push(*it);
                    }
                }

                start = stack.top();
                stack.pop();

                result.push(start);
            }

            result.print(std::cout);
        }

        void BFS(int start) {
            std::vector<bool> visited(_vertices, false);
            Queue<int> queue;
            Stack<int> result;

            queue.push(start);
            visited[start] = true;

            while (!queue.isEmpty()) {
                start = queue.top();
                queue.pop();

                result.push(start);
                std::cout << start << " ";
                std::cout << std::endl;

                for (auto it = _adj[start].begin(); it != _adj[start].end(); ++it) {
                    if (!visited[*it]) {
                        visited[*it] = true;
                        queue.push(*it);
                    }
                }
            }
        }

        void print() {
            for (int i = 0; i < 8; ++i) {
                std::cout << i << "__";
                for (auto it = _adj[i].begin(); it != _adj[i].end(); ++it) {
                    std::cout << *it << " ";
                }
                std::cout << std::endl;
            }
        }

    private methods:

        void DFSUtil(int v, std::vector<bool> &visited) {
            visited[v] = true;
            std::cout << v << std::endl;

            for (auto it = _adj[v].begin(); it != _adj[v].end(); ++it) {
                if (!visited[*it]) {
                    DFSUtil(*it, visited);
                }
            }
        }

    protected parametres:
        int _vertices;
        std::vector<int> *_adj;
    };

    class OGraph : public Graph {
    public methods:

        OGraph(const std::size_t vertices)
                : Graph(vertices) {}

        ~OGraph() {}

        virtual void addEdge(int v, int w) override {
            _adj[v].push_back(w);

            std::sort(_adj[v].begin(), _adj[v].end());
        }
    };

    class NOGraph : public Graph {
    public methods:

        NOGraph(const std::size_t vertices)
                : Graph(vertices) {}

        ~NOGraph() {}

        virtual void addEdge(int v, int w) override {
            _adj[v].push_back(w);
            _adj[w].push_back(v);

            std::sort(_adj[v].begin(), _adj[v].end());
            std::sort(_adj[w].begin(), _adj[w].end());
        }
    };
}

int main(int argc, char *argv[]) {
    srand(42);

    std::string commandLine;
    std::string type;
    int startV = 0;

    asd::Graph *g = nullptr;

    while (std::getline(std::cin, commandLine)) {
        if (commandLine.empty()) {
            continue;
        }

        std::istringstream iss(commandLine);
        std::string cmd;

        iss >> cmd;

        if (cmd == "u" || cmd == "d") {
            if (cmd == "u") {
                g = new asd::NOGraph(100);
            } else {
                g = new asd::OGraph(100);
            }

            iss >> cmd;
            startV = std::stoi(cmd);

            iss >> type;
        } else {
            const int v = std::stoi(cmd);
            iss >> cmd;
            const int w = std::stoi(cmd);

            g->addEdge(v, w);
        }
    }

    if (g != nullptr) {
        if (type == "d") {
            g->DFSRecursive(startV);
        } else {
            g->BFS(startV);
        }
    }

    delete g;

    return 0;
}
