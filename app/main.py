from flask import Flask,render_template

#Flaskオブジェクトの生成
#app = Flask(__name__, static_folder='./static/', template_folder='./templates/img') 
# app = Flask(__name__, static_folder='./templates/img/')    
app = Flask(__name__, static_folder='./static')

#「/」へアクセスがあった場合に、「index.html」を返す
@app.route("/")
def index():
  return render_template("index.html")


#おまじない
if __name__ == "__main__":
  app.run(debug=True, host='0.0.0.0', port=80)