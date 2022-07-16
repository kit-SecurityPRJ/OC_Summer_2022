from flask import Flask,render_template

app = Flask(__name__, static_folder='./static')

@app.route("/")
def index():
  return render_template("index.html")


#おまじない
if __name__ == "__main__":
  app.run(debug=True, host='0.0.0.0', port=80)