package main

type Book struct {
	cover Cover
}

type Cover struct {
	frontDesign Design
	backDesign  Design
}

type Design struct {
}

// 目的がないと大変
// 自分が印刷屋で、そのシステムの視点からの本という概念をコード化するなど
// 目的、視点、切り口がないと物事をみることはできない
// 本の機能的な部分に注目するのか、物質的な側面なのか、それらを抽出することがどう目的に沿うのかを考える必要がある
