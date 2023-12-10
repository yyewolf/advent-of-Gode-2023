package main

var input = []byte(`7.77F7F|-F.J-J7-LF|-7.FFL7F-L-7--7-JF-7F.LL.7-|FFF7..F-7-J777FF.77.L-FL-7-FF77-L7-F-F--FJFF|-F77F-7F7-.L-FFL-|-7-LJ77F7-F-FJ77.77J.J77F-L77.
F-F-J|FL-J7-L|.L|FJ|LF-7JL|J.|J.||LLJJLJ-.L7F7|L7|L7-L7J.|LJ-JJ-JJ-|F77.|.LLFJLLJ-JFJJF|-F-J7.LL7.J7J-7.7-|.LJFF.|-J7L--|..|-F--J||LF-|7.|-.
FF-JF77.L7--7F--JLF|LL-L-.7-|J-L7|7F|.J.L7.|.-7L-77J.L|.L--7-|..FL7LFJLF|7JLL|||J..|7|7.|L|F|7JFL-JF7-F-|.J-7F7..FJ|||L7F77.FL-|77JL|-FF-||7
|.||L7JFLF7L7.|J|LLJ7JF-JL-7|7-FL7L|J7JF--.|7||-FFJL|F|F|77F-|-|7L--|FLJ|L7.L|-J--.J-|-F|.LJ.-7LLJ7L7F|-F-|7LJ.FF|F||F|.||LF-L-|JL7FL.LLF-77
7F|-|L7JF|7JL.F-7.F-J-FJLJFLJ7J||L-|L--J|J--J|J|F77LF---7-F7F7FF7F|FJ7JLJ.J7F7LF|L7LF.|.LF-77|F-J|JLF7F777F-|J7FFL7|F7LJ777.-7F|.FJJ|7FL||||
L-JJJLJ-FJ-7FJ7FJ-F-7.L7F|F|7F7LJ7FL7|--|F7J.777||F7L-7FJFJ|||F7JF-J|77|L|JFLJ.FJF|7LJ-.F-JFF-J..|77|LJL77J.|LJ7.L--.L7--F-J||J-|L|JF-77LL7J
LL7F--7FJ.F--L77.|FF|-LJJF7-JFJLF77F7J|7FF7F7F7FJ||L7FJ|FJFJ|LJL77-LLJL7.|-FJ.|.FL|LFJLFL7F|..FFF7J7L-7FJL|-L7LF7FLJ7FF-.JJ7FJ7JF-J-J7LJFF|J
F.F|FJ--7.F--7|F--||..FJJ7JJLL7FFL77.|L|JF7|||||FJ|FJL7|L7L7|F--J|||LFJ77F-|FL7FLJL|J7.|-|J..|7FJ|.|LFJ|.L||-7FJ|L7L-|.|..|LF7FFJ7JFF-J7-J.J
--LJ|||.L-F-LFLJLLJ--7--LJ-.LFLFJLJ7FF.FFJ||LJLJ|FJL--JL7L7|||F7-LFJ-JF|JJFF7JL--F|F7.-J|.|L-F7L7L-7FL7L77L--LL|JL7L7.777.|JL|JJ.|F-J.L7-LL.
L7.L|-7-L-77-JFL.|JF-J|J|LFFF-J|L-JFF--7L7|L-7F-JL----7FJFJLJLJL7JJLF-JJFLFJ|-|7FF-JL7FF7-F-L||7L-7|F7|FJF77F7|L|-J7L--77LF...FLLL77|F7||J.J
FJJ.|FJ|JFL77-LLF7FJ-F-.7F7--|FJJ-FFJF7L-JL7L||F-7F7F-J|FJF-----J7.7L7F-7-L7L7F-7|F7FJ7FF7F7|||JFFJ||LJL7||F7F-7|.F---FF|JF-J77F|.JFJ--L--J|
|-|L7-7J77JLJF|.L.FL7.FJ7F7J7FF7.FFL-JL7F-7L7|||FJ||L-7|L7L-7F7F7--|7L--FF-L7|L7|||||.FFJ|||FJ|JFJFJL--7|||||7|LF7-JFFF7JFJJ.7-J|-|FL77||-7F
F---|LJLF-7-JJ-JJ|LF7FL7|F7-|L|L7F7F7|LLJLL7||||L7||F7||L|F-J|||||F|-7-F--7F||FJ|||LJF-JFJ||L7L7L7L7F7FJ||||L7F7||FLFF|L7JJ|7L|.--FJ--J--.7-
LJJ.L7|7LFJ|FLJ7L|-JJ7-FL7LF--JFJ|||L7.FF7-|||||FJ|||LJL7|L--JLJL-7JL|J|F7L-J|L7LJL7JL7FJFJ|FJFJFJFJ|||FJ|||FJ|LJL77JF|FJ|F|.7|7--J7FL7.LFJ|
-JF7.L--|L-J7F-J|L7J|FF7L7.L--7|FJ||FJF7||FJLJ|||7||L--7LJF7F-7F--JJFF7LJL--7|FJF--J|FJL7L7|L7|FL7L7|LJL-JLJL7L--7L7-F||7||LJL||FJJ|-|FF.--J
|-||7-|7|7.L|JF7|J|L7L-F-F--7FJ||FJ||7||||L7F7LJL7|L---JF-JLJ-LJ|F--7|L7FF7FJLJFJF-7-L-7|FJL-JL7FJFJL--7F----JF--JFJF-JL--7F|-JJJJFJ.F7JFLJ.
F7.F--F-LL77.7L-J.FF77-7FL-7|L7|||-||FJ|||JLJL7F-JL----7L--7F7F--JF-JL7L7|||F--JFJFJFF7||L---7FJL7||LF-JL--7F7L--7|-L7F---J---L77F|FFJ.|L|-7
LJ-|7-7.FJL7..777.FJ.77|F--JL-JLJ|FJ|L7LJL7-F7|L-7LF7F7L7F-J|LJF--JJF7L7||||L-7FJFJ.FJ||L-7F7||F7|L7FJF----J|L-7FJ|F7||F-7.|F|L|-|77.JF--L-7
LF-LFFJ-7-LLF77JLFFJ7JLL|F----7F-JL7|FL--7L7|||F-JFJLJL-JL7J|F-JF-7FJ|FJ||||F-JL7L-7L7|L-7||LJ||||FJL7L-7F--JF-JL7LJLJLJFJ-7F|7J.L7JFJ-|7FJ|
-J.|||F-77J||FJL7LF7J.|LLJF7F-J|F-7|L7F7FJFJ||||F7L-----7FJFJL-7L7||FJ|FJ||||F7FJF-JFJL7.||L-7||||L7L|F-J|F--JF-7L7F-7F-J|J||L7L|F-.J-FJFJ.7
L|-L|JF77..LLL--F-J|L-F7F7||L-7LJFJ|FJ||L7L7||||||F7|F7.||FJF7FJFJ|||FJ|FJ|||||L7L-7L7FJFJL--JLJ||FJFJL-7||.F7|FJFJL7|L-7JJLJ-|7||L7|FJJ-L-L
F77.7FJJJF7.L|J.|F7L7L|||LJ|F7L-7|FJ|FJL7|FJ|||LJ||L7||FJ||FJLJFJFJ||L7|L7LJ|||FJF-JFJL7|F---7F7LJL7|F7FJ|L7|LJL-JF-JL7FJ|FFJFL|JLF|F|LJ|L||
LJL-J-J..F7FF-.FLJ|FJ.||L-7LJL--JLJFJL7FJ||FJ|L-7LJFJ|||FJ|L7F7|FJFJ|FJL7L-7||LJFJF7L7FJLJF-7LJL--7|||LJ.L7LJF----JF-7LJLFJ.LF-|JJL.-LJ.FJL7
F||LJ.7FJ-|7|JF7LFJL7FJ|F7L--7F-7F7L--JL-JLJFJF7|F-JJ|||L7|FJ|||L7L7||F-JF-JLJF7|L||FJL-7FJFJF7F--JLJL7F7LL7FJF7F7FJFJF-77-||.F|7-.FLF-|L-77
J----7.--77FF-JL-JF-JL7LJ|7F7LJFJ|L---7F----JFJ||L-7FJ||FJ|L7||L7|FJ|||F7L--7FJLJFJ||F--JL7L7||L--7F-7LJ|F-J|FJLJLJFJJL7L--7-FJL|-LF-JFJFFJ|
L-|-L7.JJF-7L----7|F77L-7L7||F7L-JF---JL7-F-7L7LJF7|L7||L-JFJ||FJ|L7LJLJL7F-J|F7|L7|||F7-FJFJ||F-7|L7L--JL-7|L7F--7L-7FJF--J-F7.|FJ.|F7.F|-7
|F-JL77FFL7|F7F-7|||L7F7L7|||||F7.L----7L7L7L7|F7|LJFJ||F--JFJ|L7L-JF7F7FJL-7|||F7|||||L7L7||||L7LJFJF-7F7FJL-JL-7L--J|FJF7.FJL-7|.J-JJFLL-J
F-JFFJFFF7|||LJFJ||L7|||F|LJ||LJL7JF7F7|FJFJFJLJ||F7L7LJ|F7.L7L7L7F-JLJLJF-7|||LJLJLJ|L7|FJL7|L7L-7L-JFJ||L7F----JF-7F|L-JL7L7F-J-JJJ7F7.L-J
|J.||-F-JLJ|L-7L-JL-JLJL7L7FJL7F7L-J||||L7|FJF--J||L7L-7||L7FJFJFJ|F7F7F7L7LJ|L-7F--7L7|LJF-JL7|F7|F--JFJL-JL--7F7|FJFJF---JFJL7JJ.JFFJ-7J7F
F..|J-L---7|F7L--7F-7F-7L7|L-7LJL--7LJ||FJ|L7L--7||FJF-J|L7||FJJL7LJLJ||L7L-7|FFJL-7L-JL-7|F7.|LJ||||F7L--7F--7LJLJL-JFJF--7L7FJJF|.FJ|J|7F|
FF7.|FJ7F7|||L---JL7LJ7|FJL-7L7|F-7L7FJ||FJFJ7F-JLJ|LL7FJFJLJL7F-JF---JL7L7FJL7|F7FJ|F-7FJLJL7L-7||L7|L7F7||F7L----7F-JJ|F-JFJL-7F77|LJ-J-J7
|L--|7|L|LJLJF7F7F7|F--J|7F7L7L7|FJFJL7||L7L-7L-7F-JF-JL7L7F--JL-7L--7F7L7LJF-J||LJF7L7|L-7F-JLFJLJFJ|FJ|LJLJL7-F77LJJF-JL7FJF--J|L7J|L-J.J.
|LL7|JL-L7F7FJ||||||L7F7L-JL7L7|||FJF-J||JL7FJ-FJ|F7L--7|FJL7F7F7|F--J||7L-7L7FJL-7||FJL7-|L7F7L7F7|.||-L-7F--JFJL-7-FJF-7||FJF7J|FJ|F7|F7-|
||F-JFJF7|||L7||||LJ7||L---7L-JLJ|L7|F7|L7FJ|F7|FJ||F7FJ||F-J|LJ|||F7FJ|F7|L7||F--J|LJF-JFJFJ||FJ|LJFJL77FJL---JF--JFJFJFLJ||-|L7||F--77||F|
F|7JFLFJLJ|L7|LJLJ-F7LJF7F7L-7F-7L7|||||FJ|FJ||||-|||||FJ||F7|F-J||||L7|||F7||||F7-L7FJF7L7|J|||FJF7L7FJFJF----7L--7|FJF77FJL7L7||||F-J7-FF7
L7J-7-L-7FJJLJF7LF-JL7FJLJL7FJ|FL7||LJLJ|FJL7||||FJLJ||L7|LJ||L7FJ|||FJ||||||||LJL7FJ|FJ|FJ|FJLJL7||FJ|FJFJLF7LL---J|L7||FJF-JFJLJLJL7JJ|JL|
7LLJ|7.FJL-7|F|L-JF-7|L---7|L7|F-JLJF-7FJ|F-J|LJ|||F7||FLJF-JL7||FJ||L7||||||LJFF-JL7LJFLJFJL7LF-J|LJ|LJ7L--JL------JFJ||L7L--JF--7F-J.FL7F|
|7L-77-L7F-JF7L---JFJL----JL7LJL----JFJL7||F7L-7||FJLJ|F77L-7FJLJ|FJ|FJLJLJ|L--7L-7FJLF---JF-JFJF7|F--------------7F7L-JL-JF-7FJF7|L7J.|JFL7
LL7.L-7FLJF-JL---7FJF-7F---7L-7F--7F-JF7|||||F7|LJL7F7LJL7F-J|F7F|L7||-F---JF--J7FJL-7|F-7FJF-JFJ||L7F7F------7F-7|||F7F7F7|.|L-J||FJJ.-7|--
..-7LL|7|LL-----7LJFJLLJF-7L-7|L-7|L-7||LJLJ|||L-7FJ||F-7|L-7LJL7L-JLJFJF7F7L--7FJF7FJ||FJ|-L-7|7LJFJ|LJF-----J|FJ|||||||||L7|F--J|||F-JLL-7
-F7F..|J7.LF-7F7L--J-FF7L7|F-J|F-JL--J|L7F--J||F7|L7||L7|L7FJF--JF7F-7L7|LJL7F7|L7|||FJ|L7L-7FJ|F7JL-JF7L-7F7F7|L7LJLJLJ|||FJ|L--7LJ-L7FJJ77
F-7|L7|.F-F|FJ|L----7FJL-J|L-7|L------JFJL7F7|||LJFJ|L7||FJL7L---J||FJJLJF--J|LJFJ|LJ|FJFJF-J|FJ|L----JL--J|||LJFJLF7F-7|||L7L7F-JF7JJLJ7.-7
7LLJ7LJFF--J|FJF---7|L---7|F7||F7F-----JF7LJLJ|L7FL7L7|||L7J|F7F-7LJL---7|F-7L7|L7L-7||LL7L-7||7|F-7F-7F-7F|||F-JF7|||FJLJL-J||L7|||L7J.L-||
L7L|LJ-FL--7|L-JF7J|L----JLJLJLJ|L------JL---7L7|F-JFJ|||FJFJ||L7L7F-7F7|LJ|L7L-7|F-J||F7L7FJ||FJL7LJ-|L7L-JLJL--JLJLJL---7JF7L-JFJL7J-F.-L-
|LFF.|.F---JL7F-JL-JF----------7|F7F-7F------JF||L7FJFJ|LJFJFJ|FJFJL7|||L--7FJF7|||F7LJ|L-JL7|||F-JJF7L-JF-7F-7F7F7F7F7F7FJFJL---JF-JJ.F7.|J
F--F77FL----7|L-----JF---------JLJLJL|L-------7||FJL7L7L7FJFJ-LJJ|F7|LJ|F-7|L7||||LJL-7L7F7FJ|||L7F-JL---J|LJLLJ||LJ||LJLJ7L7F---7L7JF-J|-77
|J-FF-7J.F--JL--7F7F7L7F7F7F7F7F----7|F7F7F--7|LJ|F7L7L7|L7L--7F7LJ|L-7||FJ|-|||||F7F-JL||||FJ|L7|L--7F7F-7F---7LJF7LJJF7F7FJ|F-7L-JFJF-J--J
L|LLL7|FFL-----7|||||FLJLJLJ||||F---J||LJ|L-7|L-7|||FJ7|L7|F7FJ|L7FJF7||||FJFJ|LJ||||F--J|LJL-JLLJF--J|LJFJ|F7FJF-JL--7|LJLJFJ|FJ|JL|FJ7.L|J
LL7-L||F7F7F-7FJLJLJL-7F7F7FJ|LJL7F7|LJF7L--JL--JLJ||F-JFJ|||L-JFJL7|LJLJLJFJFJF-J|||L--7|7F7.F--7L---JF7L-J||L7L7F---J|F---J|||F777||F7F7|J
.LFJL|LJLJ|L7LJF7F---7||||LJFJF7LLJL7F-JL---7FF7F-7LJL7FJFJ||F-7L7FJL---7F7L7L7|F7|||F--JL7|L7L-7||F77FJL---JL-JFJL7F7FJ|F7|F7|||L7FJLJLJL-7
--L7.L7F-7L-JF7|||F77||||L--J|||F7F-JL7F---7L-JLJFJFF-J|FJFJ||-L-JL-7F--J|||L7|LJ||LJL7F7FJL7|-FJL-JL7L----7-F-7|F-J||L7||L-J||||FJ|F---7F7|
|FLJL7LJFL---J||||||FJLJ|F----JLJ|L---JL--7L7F7F7L7FJF7|L7|FJL--7JF7|L---JL7FJL7FJ|F-7||||F7|L7|F-7F7L----7L7L7|||.FJL7|||F--J|||L-JL--7||LJ
FF.FFLFF-----7LJLJ|LJF-7|L----7F7L-------7L7||LJL7|L-JLJ7||L7F-7|FJ||F7F7F7|L7FJL-JL7|||LJ||L7LJL7LJL--7F7L7L7|||L-JF7LJLJ|F--JLJF7F7F7|||J7
L7.-|-FJF---7L----JF7|FJ|F7F--J|L-------7L-JLJF7FLJF-----J|FJ|J|LJFJ||||||LJFJL--7F-JLJ|F-JL-JF-7L7FF--J|L7L7LJLJF-7|L----J|F-7F-JLJ||LJLJ-7
FJJ.F.|FJF--JF-7F--JLJL7LJLJF7FJJF------JF--7FJL7F7L--7F7FJL7L7L7FJJLJ|||L7.L7F7FJ|F---JL----7|7L7|FJF--JFJ.L7F--J7|L------JL7|L---7LJJLL|-|
L7LF|-LJ|L-7FJ-LJLF----JF-7FJ||F7L----7F7L-7LJF7LJL--7|||L7L|FJFJL---7LJL7|F-J|||FLJF-7F-----JL-7||L7L7F-JF--J|F7F7L7F-7F7F-7||F7F-J.F7.JJLJ
LJ-J|LL-F--JL7F--7L----7|J|L7LJ|L----7LJL--JF-JL-----JLJL7|FJL7L7F7F7L7-FJ|L7FJ|L--7L7LJF--7F7F7||||L-JL-7L-7FJ|||L7|L7||||L||LJ|L7JF|L77.|7
.|..LFJ.L---7|L-7L-----J|FJFJF7|F---7L-----7|FF-----7F7F7LJL--JFJ||||FJLL7L7LJFJF--JFL-7|F-J|||||||F-----JF7|L-JLJFJ|FJ||||FJL-7L-JF-JFJJJ-L
-LJFLLF--LF-J|F-JF7F7F7FJL7|FJ|||F--JF7F--7LJFJF----J|||L7F--7FJFJ||LJ||LL-J|LL-JF7-F--J||F-J||LJLJL------J||F7F7FJ.LJ|||||L7F7L-7L|F-JJ|LFJ
L|-F77||..L--JL--JLJLJLJ.LLJL7LJ|L---JLJF-JF-JFJFF7F7||L7||F-J|FJ7||JJ|7-LJL|J7F-JL-JF7FJLJF7LJLF7F7F------JLJ|||L----7||||FJ||F7L-JL-7|L77J
F|-JLLJ-7.F7F7F-------7F7F--7L7FJF------JF-JF-JF7||||||FJLJL-7LJF7||7F|-7JJ.|.FJF-7F7|LJ|F-J|F--J||||F------7FLJL----7|LJLJL-JLJL7F-7FJ77|L-
LFJ-7J|.LFJ|||L------7LJ|L-7|.LJ|L-----7FJF7L--J||||||LJF-7F-JF-J|LJ7|||.|LFFFJFJJLJLJF-7L-7|L7F7|||LJF77F--JF77F7F7.||FF-------7|L7||||L77|
J.L.7-|7.L7LJL7F7F7F7|F-JF7||F--------7LJFJL7F7FJ|||||F7L7|L--JF7L-7JLJ7.J.FJ|FJFFF7F7|FJF7||JLJ|||L7FJ|FJF--JL-JLJL7LJFJF-7F-7FJL7||L77L7J7
.|.F||LF7LL--7LJLJLJLJL--J|||L-------7|F7L-7LJLJ-|LJLJ||FJL--7FJL7FJJ.|77.F--||--FJLJLJL-JLJL7F7||L7|L7|L7L7F---7F--JF7|FJLLJ|||JJLJL-JJL|.L
F|.F7|J|FL|F7L---7F7F---7FJ||F7F7F---JLJ|F-JF7F--JF7F-J|L---7|L-7||7.-JF---JLLJ7LL-7F-7F7F--7LJLJL-JL-JL-J-|L--7|L---J||L----7LJL7.FLJ|-L77|
||-LJ-7|LJFJL----J|||F--JL-JLJLJ|L---7F7|L--JLJF-7||L7FL---7|L--JLJ-|FJ||FLJ.J.77FFJ|FLJLJF7L-7F7F7F7F-7F--JF7FJL7F7F-J|F----JF|FF77JL7.LJ77
LJ.LL7JJ77L7F----7|||L7F-7F7F-7FJF7F7LJ|L7F-7F7L7LJL-JF7F7|LJ|F7F77.L|L|77.|FF-F--JFJF7F--JL--J|LJ|||||LJF-7|LJF7LJLJFFJL---7F7F7|L-7-777JL7
FJ-JJLJ.77FJ|F---J|||.|L7||||FJL-J|||F-JF|||||L7L-----JLJL-7F-JLJL--7J|LJLF|-|||F7FJFJLJF7F7F7FJF-J|||F--JFJL--JL7F7F7|F----J|LJ||F-J-LL|.7J
F7LFF7|F--JFJ|F7F7|||FJFJLJLJL-7F7LJ|L--7||FJL7|F---7F--7F7LJF-7F7F-J.7.L-7J7|-LJLJ|L--7|LJLJLJ.L--J||L--7L-----7LJLJ||L7LF-7|F-J|L-7FFJJFFJ
|7.FL|7L7F-JFJ|LJ||||L7|F------J||F7L---J|||F7||L--7|L-7|||F-J.||LJ|LJ-L--FJFJ7.||LFFF-JL-77F--7FF-7||7F7L-7F--7L-7F-JL7L-JFJ||F7|F-J7||J|J7
J--7|LJ|LJ|LL7|LFJ|||FJ|L7F7F--7|LJL----7LJLJ|||F--J|F7||||L--7LJ-|7|-|7FL|||L-L77-FFJF---JFJF7L-JFJLJFJL--J|F-JF7LJ7F7|F--JFJLJLJL7LF777L-J
|J|.F-F-|7F.LLJFJFJLJL-JFLJLJJFJ|F--7F7FJ7F-7|||L--7||||||L7F-JF-7.7J-F|...J77LFLFF7L7|F--7L7|L--7|F7FJF---7||LFJ|F7FJLJ|F7-|F7F-7FJ-F77-|LL
LFFF7LF-FJ7.FLFL-JF-7F-7F7F7F7L-JL-7LJ|L--JFJ||L7F7|||||LJFJL--JFJ7|..FJ.|-|JJLF-FJL-J||F-J7LJJF-JLJLJFJF--J|L7|FJ|||F-7LJL-J|LJJLJF-J|7|L.J
FFJ||--7|F7-|.FF-7L7|L7LJ|||||F7F--JF7L----JL||LLJLJLJ||F7L7F7F7|J7JL-||F7|J|7F|.L7F7FJ||F7F7F7L7F--7FJLL--7|FJ|L-JLJL7|F7F-7L---7FJF7L77FJ.
J7-JJ-LFJJJ.LF7L7L7|L7L-7LJLJLJ|L---JL------7LJF7F7F7LLJ|L7LJ||LJ||.|.L|J|7-JF-LL7||||FJLJ||||L-J|F-J|F----J||FJF-7F--JLJ|L7|F--7LJFJ|FJLJJ7
L||L.F.|7J-F.||7L7||FJF7L-----7|F7F-7F------JF-JLJ||L7F7|FJF7LJF7F7-F-FJL|-7J|J.LFLJLJL7F7LJ|L7F7|L7FJL7F---JLJFJJ|L----7L-J|L-7L7FJFLJJ-|.F
LF-7FF-J7-FJF|L--JLJL7|L-77F7FJLJ|L7LJF7F-7F7|F7F-J|FJ||||FJL--JLJL7J..F7J|-.L77L|J|F7.LJL-7||LJLJJLJF7LJF-7F7FJF7L7F7F7L--7|F7|JLJJLJ|...FF
|L-J-7..--F7FJF--7F-7|L-7L-JLJF-7L7L--JLJFJ||LJ|L--JL-JLJLJF-7F7F--J||-.F7J.7-J|.LJF||F-7F7||F------7||F-JFJ||L-JL7LJ||L---JLJ||.||.|JJL.--.
F-F|-|F7.F|LJFJF7LJFJL--JF7F-7|JL7L-----7|FJL7FJF--7F-7F7F7|FJ|LJ-|7LJL.LJ.FLLFF7JFFJ|L7|||||L-----7|||L-7L7||F7F7L-7|L--7F7F7||F777..LLFJ|7
JFF--|F--L|F7|FJL--JF7F--JLJ|LJF7L-----7LJL-7|L7|F7LJ-LJLJLJL7|7JFJLF7JFLL-F7FF||F7L7|FJLJLJ|F-----JLJL-7L7LJ||LJL--JL--7LJLJ|LJ|L-7--F7J7L|
|7J|.L-7F7LJLJL----7|||F-------JL-----7L7F7FJL-JLJ|F7FF----7LLJ7LJJL-J77.FFF7-FJLJL7|LJF---7|L7F--7F7F-7L-JF7|L----7F7F7|F7F-JF-JF-J7J|F.JFJ
J|FL7|.7JF7||F-7F7L||||L--------7F---7L-J|LJF7F7F7LJL7L7F--J77|7-|JLF7-F-7F||FL--7FJL--JF--JL7LJF7LJLJFJF7-|||F-7F-J||||LJ|L--JF-J7L77LJFL|.
|L7FF--J-J7|FL7LJ|FJ|LJF--------J|F--JF7F|F7|||LJ|F7FJFJL7LF7F7.F-7.FLJL7|7||F77FJL--7JFJF--7L--JL7F7FJFJL7||LJ-LJLFJLJL7LL----JF777LFJL-JF.
F-J-J|J.|-F7F7L-7||FJ.FJF7F7F--7FJL---JL-J|LJ|L7FJ||L-JF7L-JLJL-7|F7JF77|L-JLJL7|F7F7L-JFJF-JF---7LJLJFJF-J|L------JF--7L----7FFJL7--..|.FL|
L7J-|77FLFJLJL-7||LJF7L-JLJ||F7LJ..F7F----JF7L-JL-J|F7FJ|F7F----JF-7-FLFL-----7|LJLJL7F-JFJF7|F--JF7F-J.L--JF-------JF7|F----JFJF-J||77F|--J
|||F.F-JJ|F---7LJL--JL----7LJ|L--7FJ|L---7FJL-----7LJ||LLJ|L-7LF-JFJF--F---7F-JL7F7F7LJF7L7|LJL---JLJF77F---J|F7LF-7F||LJ7F---JFJ7FLF7F7.7.7
F77JL.|FF||F-7L-----7F--7FJF7L--7|L7|F7F7LJF-----7L-7|L7F7|F-J.L-7|F7-LL--7|L-7FJ|LJL-7|L7LJF7F7F7J.FJL7L-----JL7L7L-JL--7|F-7FJLFJ|FJFL-JFJ
||L--7-JLLJL7L-----7LJLFJL-J|F7FJL-JLJ|||F7L----7L--JL-J|LJ|F77F7||||F---7|L7FJL7L---7||FJF-JLJLJL7FJF7L--------JFJF7F7F7LJL7||F7J||J.||7JFJ
|-7LL|JF--F-JF----7L-7FJF--7LJLJF7F--7LJ||L-7F-7L-7F7F77|F-J|L7|LJ|||L--7||FJL7FJF--7|LJL7L-7F-7F7LJFJL-----7F7F-JFJLJLJL---J|LJL7F7LFFF-F-7
|7|L||F|L||F7|F---JF7LJFJF-JF7F-JLJF7L-7|L-7|L7|F7||LJL7LJF7|FJ|F7||L7JFJ|||JFJL7L-7|L-7FJF7LJ7LJ|F7L7F-----J|||F-JJ7F7F----7L7F-JF7.|LFJ|||
--LFFJ7J7FLJ|||F7F7||F7|FJF-JLJF---JL7FJL--JL7|||LJ|F-7L7FJ||L7LJ||L7L-JFJ|L7|F7L7FJ|F7||-||F----J||FJL7F7F7FJLJL--7FJ||F---J-LJF7||-7FLFJJ.
LF7||L-7-FJFJ|LJ||||||LJL-JFF--JF--7.LJF----7LJ||F-J|7L-J|FJL7L7FJ|J|F--JFJFJLJ|FJL7||||L7|LJF--7FJ||F7LJLJLJF----7LJFJ|L-----7FJ|||77LF77L.
||LJJ7.J.L-L-JF7LJLJLJF-----JF--J|FJF--JF7F7|F7LJL-7L7F7||L-7|FJ|FJFJ|F7-L7|LF-JL--JLJ||FJL--JF7LJ|LJ||F7F7F7L---7|F-JFJF-7F--JL7LJL-77||7JF
L7LJ7-|-FF|JF-JL---7F7L--7F--JF7F7L-JF--JLJ|LJL7LF7|FJ|L7L7FJ||FJL7L7LJL-7|L7L--7F-7F-J|L7F7F7|L7F---JLJLJLJL-7F7|||F7L7|F|L---7L7F7FJJFJ|-7
|--L-F---7JLL----7FJ||7F-J|F--JLJL---JF7F7||F7FJFJ||L7|FJFJ|7||L-7L7|F---JL7L--7LJFJL-7|FJ|||||FJLS---------7|||||LJ|L-JL7|F---JFJ||L77L7-.|
|LF..L|LFL7.LF---JL-JL7L7FJL---------7|LJL7LJLJ7|FJ|FJ|L7L7|FJ|JFJFJ|L7FF7FJF--J-FJF--J||FJLJLJ|F7F77F7F---7L7||LJFFJF7F-J||F7-FJFJL-J7..LF7
|.|.FJ|L|L|7FL7F-----7L7LJF------7|F7LJF-7L7F--7|L7|L7L7|FJ|L7L7L7L7|FJFJ|L7|F77FJFJF-7|||F---7|||||FJLJF-7L-JLJF7FJFJ||F7LJ|L7L7L7F7F--7-F7
|7FFJ-L-J|F-JLLJF7FF7|FJF7L---7F7L-JL--JJL-JL7FJ|FJ|FJFJ|L7L7L7|.|FJ||JL7|FJ|||FJFJFJFJ||LJF7|||||||L---JFJ-F--7|LJFJ|LJ||F7|FJFJFJ|LJF7L7J|
.FJ|LF|-|7.||7-FJL-JLJL7||F--7LJL--7F-----7F7||FJL7||7L7|-L7L7||FJL-JL7FJLJFJ|||FJ7L7L7|L7FJL7|||LJL7F7F-JF7|F-JL7FJF7F7|||||L-JFJFJF7|L-JL-
-L-7-LL-L--7LF-JF7F7F-7|||L-7L-----J|F----J|LJ|L-7|||F-JL7J|FJ||L7F---JL--7|FJ||L-7.L7||FJL-7||||F--J|LJF-JLJ|F7FJL-JLJLJLJLJF-7L-JFJLJJ.|7.
|-FJ|.J7FL-|-L-7|||LJ||||L7FL-------JL-7.F7L-7L7-||||L-7FJFJL7||FJL7F7F-7J|||FJ|F-JF-J|||F7FJ||LJL7F7L--JF--7LJLJF7F--7F7F--7|-L7F7L-7JLFL|7
|-J-J-77FLF.F|-LJLJJF-J|L7|F7F------7F7L7|L7.L7|FJLJL7L||FJF-J||L7FJ||L7|FJ||L7||F7L-7|||||L7|L-7FJ||F---JF7L7F--J|L-7LJLJF7|L-7LJ|F7|-F-7LJ
-J.F|F-LJLF--7F7F---JF7L7||||L7F-7F7LJL-JL7L77||L-7F-JFJ|L7|F7||FJL7||FJ||FJ|FJ|LJ|F7|||||L7|L7FJL7||L----JL7LJF-7|F7L7F7F||L--JF7LJLJ|L-JLL
LF7|F-JL7|L7FJ||L----JL7LJLJL-J|L||L7F7F-7L7|FJL7FJL7FL7L7||||||L-7LJ||FJ|L7||FJF-J|||LJ||FJL7LJF-J||F-7F--7L--JFJLJL-J|L-JL---7|L7-||JFLJL|
L|JF7-7---FJ|FJ|F7F7|F7L7F7F-7FJFJ|FJ|||FJ7||L7FJ|F-JF7|FJ||||||F7L-7|||7L7|||L7|F7||L7FJ||F7|F-JF7||L7|L-7L---7|F--7F7|F------J|FJF7--J7-F|
7-FJL7LF--L7|L7LJLJL-JL-J|||FJ|JL-JL7|||L7FJL7||FJ|F7|||L7||||||||F7|||L7FJ||L7||||||FJ|FJ|||||F7|||L7||F7L---7|LJF-J|LJL-----7FJL-J|.-J|.F7
L-7---FJLFFJ|FL---7F-7F7FJLJL-JF----J|LJFJL7LLJ|L7LJ||||FJ|||LJLJ|||||L7LJL|L7||||LJ|L7|L7|||||||||L7|||||F7F-J|F7|F7|F-7F----J|F---J7J7F7F7
FF--JF-7.LL7L-----J|FJ|LJF--7F-JF--7FJF-JF-JF--JFJF-J|LJL7||L--7FJ|||L-JF--JFJ|||L7FJFJL7|||||LJ||L-JLJ|||||L7JLJ|LJLJ|-||-F---JL---7JFL7FJJ
L7JL-JJ|7|F|F7F7F7FJL7|F7|F7LJF7|F-J|FJF7L7.L7F7|FL7FJF--J|L7F7|L7|LJF--JF7FJFLJL7||FJF-J|LJ|L-7|L---7FJ|||L7L--7L7F-7L7LJFJF7F7F---JFJL|-|J
FF7L7|.LLF-J||||||L7FJ||LJ|L7FJ||L7FJL7|L-JF-J||L-7|L7L--7|FJ||L7LJF-JF--J|L-7JF-J|LJLL-7L-7|F-J|F---JL7||L7|F--J||L7L7L--JFJLJ|L--77J.LFJ|7
J.FL7F7J|L-7|LJLJL-J|FJL-7|||L7LJJLJF7|L-7FJF7|L7FJL-J7F-J|L7|L7||FJF7L--7|F7L7|F7L7-F--JF7|||F-JL7F7LFLJL7||L--7FJFJ|L7F7FJF7FJF--J-7F7|-JJ
FLJFLJL--||||JLLF7F-J|F-7LJFJFJ.F7F-JLJF7|L7||||||F----JF-JFJL7||FJFJ|F7FJ|||FJ||L7|FJF-7|||LJL7F7LJL7F---J||F7FJL7L--7|||L-J|L7L7J|LFJ--77.
FJ.-7.|.-LLLJ.F-JLJF7|L7L--JFJF-JLJF7F7|LJ|LJ||FJ||F7F-7L77L7FJLJL7L7LJ|L7LJ||FJ|LLJL-JJLJ|L-7F||L7F7||F7F7|LJLJF7L7F-J||L7F-JFL-J.J.77F|--|
L-7FJ.F-|7J.FFL---7|LJFJF7F7L7|F7F7||||L--77FJ|L7||||L7|FJF-JL---7L-JF7L7L-7LJL7|F---7LF7FJF-JFJL7|||||||||L----J|FJL-7||FJL----7.-L-LJLLJ7|
J-777.7LL---LF----J|F-JFJLJL-JLJ||LJLJ|F--JFJFJ7||LJ|FJ||JL-7F7F7L7F-JL-JF7L-7FJ|L--7L-J|L7L7FL7FJLJLJ|||||F7F7F7|L7F-J||L-7F7F7L7.|J|JFL-JJ
LF|J7LJFJ.|7FL--7F-JL7FJF-------JL-7F-JL-7FJFJF-JL7|LJFJL-7|||LJL-JL7F--7|L7FJL7||F7L--7L-JFJF-JL---7FJ|||||LJ||LJ7|L-7||F7LJ|||FJ7F7FFJJJL7
FL7-|J.JJ7LFF---JL77|||FJF7F-7F7F-7||F-7FJ|FJFJF-7|F--JF7FJFJ|F--7F-JL7FJ|J||F7||FJL--7L7F-J7L-7F--7|L7||||L-7|L7F7L-7||||L7FJ||L--77FJL-7|F
7J.||-7J7|7|L-7F7FJF7||L7|||FJ|||FJ|||FJ|FJ|FL7|FJ||F7FJ|L7L7LJF-JL7F7|L7|FJ|||LJL-7F7L-JL---7FJL-7LJFJ|||L7FJL7LJL7.|||LJFJL-JL7F-JJL-7|LJ-
-7F-J.LF-|-L..||LJFJLJ|.LJLJL7|||L-JLJL7|L-JF7||L7|LJ|L7L7|FJF7L--7||||FJ|L7||L----J|L7F7F7F7|L7F7|F7L-J||FJ|F-JF-7L7LJL7FJ7J|JL|L--7J.|F7LL
LFJLLF-7LF7JF-LJLL|F-7L-7F---J|LJF-----JL---J||L7||F-JFJFJ|L7||F7FJ||||L7L7|||F7F---JL||||LJLJFJ||||L--7LJ|FJL7FJ|L7L-7FJL-7F77-L7F-J.77FJLL
.|7F|LJL-J|7|LLF--LJLL7FJ|F-7FJF-JF---7F7F---JL-JLJ|F7L7L-JFJ||||L7LJ||FJFJ||LJ||F7F--J||L7F--JFJ|LJF-7L-7|L7FJL--7|F7||F-7LJ|--FJ|7|F-FJ7.|
F--7LJJ.|L|.L7FL-LJ||FLJFLJFJL7|F-JF--J||L----7F---J||FJF--JFJ||L7L-7||L7|FJ|FFJ||LJF-7||FJ|F7FJFJF7L7L--JL-JL7F-7|||||||FJF-JJ.L7L7-|JL|F.F
LJ7||JF-|JFJ7L--|J||7FLLJ-LL7FJLJF7L--7|L7F---JL---7|||7L7F7|FJ|||F-J||FJ|L7L7L7|L7FJFJ|LJFJ|||-L7|L-JF7-F--7F|L7LJLJLJLJL7|J|.-LL7|7L77LL-7
F-7.-.7FL.LJF7J-|F|.J7LLL-|7LJLLFJL---J|-|L--77F7F-J|||F7LJLJL7|FJL-7||L7|LL7|FJL7|L7L7L-7L7||L7FJL---JL-JF7L7|FJ|-|||7|.LLJJFJ..F||7-FJ|.|J
L.L-LFF-FJ.-L-J.FLF-F7-.FL|F-77FJF7F7F7L7|F-7L7||L7FJ|||L-----J|L7F-J|L7|L7F|||F-J|FJFJF7L7||L-J|F7F7F--7FJ|FJ|L-7-J7JFJ-L|L-FJ-|-||--FJL-|J
.FL7.-FJ|L7.||L-JFJ|LF.77FLL7L-JFJLJLJ|FJ|L7L7LJ|FJ||LJL--7F7F7|FJL-7|FJL7L7||LJF-J|FJFJL-JLJJF-J|LJ|L-7LJFLJ7L--J||JL|L77|..|7-|.LJJFJ.|-|-
|JF-L7LF|.-7J.77.|F7F7-77F-L|F7FJF--7FJL7L7L7L7FJ|FJF-----J|||LJL7F-JLJF-JFJ||F-JF7|L7L---7F--JF7L-7|F7L7F-7J7J7|JF---JFJLFJ7F-7LL777FJ.7F-J
L-F-LF--J.JJ.FL77L-JF|FLJ||.LJLJ-|F7LJF7|J|FJL|L7||7|F-7F-7|LJF--JL--7JL7FJF||L-7||L7|F-7FJL-7FJL7FJLJL7LJFJLJ.FJ.F-|F-J77|L-JJJ-|LLLL-|-L.|
7LF7.FJ|FFJ|F|JL7--7.LJ-FL7F|J..FLJ|F7||L7LJF-JFJ|L7||FJ|FJ|F-JF-7F-7L7||L7FJ|F-J||FJ|L7|L-7FJ|F-JL--7FJF-J|J.FLJF|F-7F-7---J.||FL7L7J|L|J.|
L.-|--J|.|.7--7J7J..FJL-L-J||--FF-|LJ||L7L7FL-7L7L7|LJ|FJ|FJL-7|FJL7L-JFJFJ|FJL7FJ||FJFJL7FJ|FJL7F7F-JL7L--7|FJL-J|7.|7|JLLL-7FL---7L.---777
J-L|-|7.FF7|.L|.|J|.---|7|7F7FLL|FFJL||LL7|JJLL7|FJ|F-J|FJL-7FJ|L-7L--7|FJJ||JJLJ7LJL7L7.LJ7||7LLJ|L--7L7F7|J-|L|7.F-7||7-L77L7-J..F-FL.LJ-7
L7|F-7J7-FJ7F7LJ-F7.F|-F-77LL7LJ|LJLFJ|-LLJJL|LLJL7|L-7||F--JL7L-7|F7FJLJFLLJF--FF---JFJF---JL7-F7|F7FJJ||LJJFL.-J.LFJJJF-7L|7J|L-F-FLJ.|||.
.|-L-|L-JL.77L7|-L|-FF.||.--FF.|F-||L-JLL-LL-F7LF-J|LFJ||L77-FJF7|||LJJF77||F|J|LL--7FJJL7F-7FJFJLJ|LJF7LJ..-7JFJ-L--JJ-J-LLFJ.||FJJ.|.JLF|J
.LJL77-JF|7|7-LL.LJLLJ.LF-|-7.LLJ.F-JL|-L7L|.7||L--JJ|FJL7L-7L-J|||L--7||L-L7L77.J7F||F--J|FJL7L7F7L--JL7J77F-J||F|J||.L|J|F|7-L-JJ77J.FFFJ.
-J..L7JLFJF7-7F|77L7-77F-FJ-|.|-F.|L|L7|LJ7F7JFF|J|F-JL7LL--J|LFJ||F-7LJL7.L--.J-FFFJ||F7FJL-7L7|||F-7F-J|FF-.FLJ-J.7F7.-.F7|FJFJ7LFJF-|-JF7
J--7JL7FJF|LJ-LJL--JL|F--JJ-|FF7LLF--FJJ|F-.7|FLJ|LL--7L----7J7L7||L7|F7FJJJ|FJJ-FFL-JLJ|L7|FJFJLJ|L7LJLLFFJL|77||7.L|LJL-JLLJ7-7LJ|-J.||.||
|-L|JL||F7L777||LJ-L.FJF7J.FLF--7FLF-|JL7-LJ-|L|-JJ..FJF-7F-JJFL|||FJ||||J.-JJ|F-|JJ|FF-JFJFJFJ7.||FJ-J7|J|7F|-L-J-J-J7J.|FJLJ.FJF7..F77L7-L
|||.F.LL||-L-J7-LFF7-|--LJ.LF|JLJ7FF7L|FF-7J.LFJ-77F-|FJFJ|.|J|FLJLJLLJLJJ7|LJFLJLLJ|7L-7|-L-JJ7.FJ|-F7JL--J-J7JJLFJLL7-FFL77FLJJL7-.7J||F-.
FLJ7LLJFL-7FJF77.LL7J|F||7|-LJ7.FJ7|LF77JLJ7FF-LJ-J77LJ|L-J--7.L-L|JL.||JLJF-FJL77.-.L.LLJJ.|.F|FL7|-L7FFJ-L-J|J7FJ7JFJ7|F7L77|FF-L-7|-LFJ.F
FJ.7LL--L-7J--L-7LJJL-|-FFL-F------JF|-7-7.-FLJ.L|LL|JL|JLJJLL7L-JJJLFJ.LL-JLJL7--F--J-|.J.L-J-LFJLJ7J-7-JLLFLJ.J.LL.JJ-JLJL-J|---L7LFJLLL|.
`)

var inputTest = []byte(`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`)

var inputTest2 = []byte(`FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`)
