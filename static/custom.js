feather.replace({'aria-hidden': 'true'})

let alertList = document.querySelectorAll('.alert');
alertList.forEach(function (alert) {
  let alert1 = new bootstrap.Alert(alert);
  setTimeout(function () {
    alert1.close();
  },3000)
});
