# OC_Summer_2022
ユーザーを見つけてもらうサイドチャネル攻撃をする
その結果見つかったユーザー名を憶えておいてもらう
ユーザー名を使用してデータ検索にて購入履歴のSQLインジェクションをしてもらう？　←今はとりあえずこれ
自動でユーザーを最初に登録しといてそれ以外のユーザーの情報が見えるようにする？

SELECT hoge FROM purchase WHERE productname = "%s"
としてOR '1'='1'で全部表示出来る感じにする できるはず
guestユーザーのデータが見れるだけなのに他のも見れるようにする？

今のとこの進捗
/login
/search
のエンドポイントに関する処理を書いた

フロントがまだ
loginページ　未
    エラーメッセージを表示させるようにしたい
injectionページ Done（エンドポイントが/search/testのまま）

DB立ててない　
dml書いてない Done
dockerfile composeyml書いてない Done

時間あれば
今ある物は、バックエンドのみなので出来れば、クリックジャッキング的なものも欲しい押したらエビルポータルにサイトに繋がるような感じ←これは結構時間かかりそうなのでいけそうだったら（出来たらおもろい）

DBユーザーテーブルパスワード
user|password
----------------------
Bill|Microsoft
Satoshi|Bitcoin
John|personalcomputer
Steve|iPhone
Alan|enigma
Linus|linux
Yukihiro|ruby
Lisa|ryzen
Kevin|shimomura
Peter|spiderman
---------------------
guest|guest