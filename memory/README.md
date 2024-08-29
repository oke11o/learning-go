
Работа GC в GO, как распределяется, высвобождается, уплотняется используемая в приложении память.
Как происходит захват и освобождение переменных в лямбдах. 
Чуть подробнее узнать о строении фреймов стека горутин.

цикл статей на ardenlabs:
https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html

Хабр Go To Memory
https://habr.com/ru/companies/oleg-bunin/articles/676332/


Относительно фреймов стека горутин можно в целом затронуть тему Bound Check Elimination. В статьях по BCE часто дается инфа по устройству стека:
https://medium.com/a-journey-with-go/go-memory-safety-with-bounds-check-1397bef748b5
https://medium.com/a-journey-with-go/go-how-does-the-goroutine-stack-size-evolve-447fc02085e5


Как происходит захват и освобождение переменных в лямбдах - а вот про это накидайте материала?
https://dr-knz.net/go-calling-convention-x86-64.html#toc-entry-17 - Здесь есть абзац про deferred closures с ASM кодом.
https://habr.com/ru/companies/badoo/articles/468863/ - Про реализацию closures (внизу)