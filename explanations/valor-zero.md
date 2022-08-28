Valor zero é o valor que as variáveis declaradas, mas não inicializadas, recebem por padrão.

Valores zeros:
  - ints: 0
  - floats: 0.0
  - booleans: false
  - strings: ""
  - pointers, functions, interfaces, slices, channels, maps: nil

Por isso, é interessante utilizar o := sempre que possível, para não ser recebido de surpresa por um valor inesperado.

Use var para package-level scope.