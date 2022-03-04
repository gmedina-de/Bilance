feather.replace({'aria-hidden': 'true'})

let alertList = document.querySelectorAll('.alert');
alertList.forEach(function (alert) {
  let alert1 = new bootstrap.Alert(alert);
  setTimeout(function () {
    alert1.close();
  },3000)
});

$('table').bootstrapTable({
  onClickRow: function (row, $element, field) {
    window.location.href = window.location.href.split('?')[0] + "/edit?ID=" + row.ID
  }
})