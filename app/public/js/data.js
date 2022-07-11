async function get() {
    const textbox = document.getElementById("search");
    const URL = "http://localhost:8080/data/test?search=" + textbox.value;
    const parameter = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    }
    const result = await fetch(URL,parameter).then((respose) => {
        return respose.json();
    });
    let table = document.getElementById("tarta");

    result.forEach(list => {
        let newRow = table.insertRow();
        Object.keys(list).forEach(key => {
            let newCell = newRow.insertCell();
            let newText = document.createTextNode(list[key]);
            newCell.appendChild(newText);
            //console.log(`${key} ${list[key]}`)
        })
    })
}
