function minimizeSelection(inputString) {
    var checkBox = document.getElementById(inputString);
    var text = document.getElementById(inputString+"-response-div");
    if (checkBox.checked == true){
        text.style.display = "block";
    } else {
        text.style.display = "none";
    }
}