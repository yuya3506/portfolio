# 旅行プラン作成アプリ（仮）

## 背景
私は、大学時代から趣味で一人旅をしており(36/47都道府県制覇)、
特に、より多くの名所を効率よく回ることを重視している。
その際、交通手段は公共交通機関を用いるため、綿密なスケジュール設定が必要である。
そのため、これまで交通機関等の運行情報や各観光地の所要時間を一つずつ調べ、手作業で計画を立ててきた。
しかし、一回の旅行は１週間程度と長いため、計画を立てることが面倒に感じており、自動で中長期的な旅行計画を立ててくれるものが欲しいと感じた。

## 類似アプリ
**NAVITIME Travel**
* 自分が立てたプランを記録することが可能
* プランの提案機能はない

**ぷらる**
* 旅行計画作成機能を有する
* 二日以上の計画には未対応

**二日以上の計画を提案してくれるアプリは見当たらない**

## 要件定義
### 機能要件
#### 旅行プラン作成機能
入力： 
* ①スタート・ゴール地点
* ②旅行期間(複数日に跨ることも可)・ホテルチェックイン/アウト時刻(必須)
* ③訪れたい観光地(滞在時間入力必須)
* ④移動手段条件(ex.低価格重視、快適さ重視など)

出力：
* ①移動ルート
* ②各地点の到着・出発時間
* ③移動手段
* ④宿泊先候補(優先度低)

**プランを作成すること自体は、会員登録不要<br>
作成したプランの保存・編集等を可能にするには会員登録が必要**

#### 会員登録機能
* ユーザーIDを元に、過去に登録したプランデータの取得
#### 退会機能
* 退会処理の実施
#### 旅行プラン保存機能
* 作成したプランをDBに保存
#### 作成済み旅行プラン(マイプラン)表示機能
* 過去に作成したプランの表示
#### マイプラン編集機能
* プランの編集・更新が可能


## 技術選定

### フロントエンド
#### 言語
JavaScript
#### 選定理由
業務で用いたことがあるため

### バックエンド
#### 言語
GO言語
#### 選定理由
* Web開発に適している
* 大規模サービスの開発に用いられており、習得することで、これらの開発に携われると考えたため

### その他
Docker


## 各種設計
### API
https://travel.app.com/plan/create 旅行プラン作成画面
### DB(会員登録機能を実装する場合)
