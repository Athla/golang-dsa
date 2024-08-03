class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        op = ["+", "-", "*", "/"]
        stack = []
        for t in tokens:
            if t not in op:
                stack.append(int(t))
            else:
                res = 0
                a = stack.pop()
                b = stack.pop()
                match t:
                    case "+":
                        res += a + b
                    case "-":
                        res += b - a
                    case "*":
                        res += a * b
                    case "/":
                        res += int(b / a)

                stack.append(res)
        return stack[-1]
