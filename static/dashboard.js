feather.replace({'aria-hidden': 'true'})

document.querySelector("#logout").addEventListener("click", function (e){
  e.preventDefault();
  let request = new XMLHttpRequest();
  request.open("get", "/logout", false, "false", "false");
  request.send();
  window.location.replace("/");
});

let alertList = document.querySelectorAll('.alert');
alertList.forEach(function (alert) {
  let alert1 = new bootstrap.Alert(alert);
  setTimeout(function () {
    alert1.close();
  },2000)
});
