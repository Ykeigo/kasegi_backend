# kasegi-backend

kasegi（仮）のバックエンド


デプロイ手順
1. EC2を立ち上げる（ここまではハンズオンの資料でなんとかして）
2. ECRにログインする
sudo aws ecr get-login-password --region ap-northeast-1 | sudo docker login --username AWS --password-stdin 744778349286.dkr.ecr.ap-northeast-1.amazonaws.com
3. EC2にログインしてECRからpullする