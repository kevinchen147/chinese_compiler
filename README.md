# chinese_compiler
A toy compiler for chinese language.

## Build

**Notice**: Compilation requires Go 1.13+

`go build -o compiler ./main/parser.go`

## Help

`compiler -h`

## Usage

`compiler --in ./example/for.txt`

You must specify the input file.

Some examples can be found in folder `example`.

## Example

### Code Input

```
字符串型 今天。
今天 为 今天星期几（）。
如果（今天 等于 “星期一”）【
	上程序设计语言原理课（）。
】另外如果（今天 等于 “星期日”）【
	写程序语言设计原理作业（）。
】否则【
	干别的事情（）。
】

```

### AST Output

```
「第 1 行｜第 3 列」 -> *ast.Block: 
    「第 1 行｜第 3 列」 -> *ast.Decl:
    「第 2 行｜第 3 列」 -> *ast.Decl:
    「第 3 行｜第 11 列」 -> *ast.ExprStmt:
        「第 3 行｜第 5 列」 -> *ast.assign: 常规赋值
            #左值:
                「第 3 行｜第 0 列」 -> *ast.IdentExpr: 建校年份
            #右值:
                「第 3 行｜第 7 列」 -> *ast.IntLiteral: 1952
    「第 4 行｜第 6 列」 -> *ast.ExprStmt:
        「第 4 行｜第 3 列」 -> *ast.assign: 常规赋值
            #左值:
                「第 4 行｜第 0 列」 -> *ast.IdentExpr: 校龄
            #右值:
                「第 4 行｜第 5 列」 -> *ast.IntLiteral: 0
    「第 5 行｜第 0 列」 -> *ast.For:
        #初始化:
            「第 5 行｜第 6 列」 -> *ast.assign: 常规赋值
                #左值:
                    「第 5 行｜第 3 列」 -> *ast.IdentExpr: 今年
                #右值:
                    「第 5 行｜第 8 列」 -> *ast.IntLiteral: 1952
        #条件:
            「第 5 行｜第 16 列」 -> *ast.Binary: 小于等于
                #左值:
                    「第 5 行｜第 13 列」 -> *ast.IdentExpr: 今年
                #右值:
                    「第 5 行｜第 21 列」 -> *ast.IntLiteral: 2020
        #步长:
            「第 5 行｜第 30 列」 -> *ast.IncDec: 加一
                「第 5 行｜第 27 列」 -> *ast.IdentExpr: 今年
        #循环体:
            「第 5 行｜第 34 列」 -> *ast.Block:
                「第 6 行｜第 12 列」 -> *ast.ExprStmt:
                    「第 6 行｜第 4 列」 -> *ast.assign: 常规赋值
                        #左值:
                            「第 6 行｜第 1 列」 -> *ast.IdentExpr: 校龄
                        #右值:
                            「第 6 行｜第 9 列」 -> *ast.Binary: 加
                                #左值:
                                    「第 6 行｜第 6 列」 -> *ast.IdentExpr: 校龄
                                #右值:
                                    「第 6 行｜第 11 列」 -> *ast.IntLiteral: 1

```

