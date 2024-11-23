# Todo App

このプロジェクトは、シンプルな Todo アプリケーションです。Go バックエンド、React フロントエンド、そして PostgreSQL データベースを使用しています。

## 開発環境セットアップ手順

Docker を使用して、簡単に開発環境を構築できます。

### 必要なツール

- [Docker](https://www.docker.com/) (最新バージョンを推奨)
- [Git](https://git-scm.com/) (リポジトリをクローンするため)

### セットアップ手順

1. **リポジトリのクローン**
   ```bash
   git clone https://github.com/<あなたのユーザー名>/todo-app.git
   cd todo-app
   ```

## プロジェクト概要

このアプリケーションは以下の機能を提供します:

- タスクの作成、編集、削除
- タスクの完了状態のトグル
- ユーザーフレンドリーなインターフェース

## 環境変数

以下の環境変数を使用します:

- `DATABASE_HOST`: データベースのホスト名
- `DATABASE_USER`: データベースのユーザー名
- `DATABASE_PASSWORD`: データベースのパスワード
- `DATABASE_NAME`: データベースの名前

### 動作確認方法

- **バックエンド API**:
  - URL: `http://localhost:8080/api/tasks`
  - メソッド: `GET`, `POST`, `PUT`, `DELETE`
  - JSON 形式でタスクを管理。
- **フロントエンド**:
  - URL: `http://localhost:3000`
  - フロントエンド画面でタスクの管理が可能。

## 今後の開発予定
- ユーザー認証機能の追加
- タスクの検索/フィルタリング機能
- ダークモード対応
