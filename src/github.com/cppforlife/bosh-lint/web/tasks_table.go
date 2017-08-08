package web

const tasksTable string = `
<table id="task-tmpl" class="tmpl">
  <tr>
    <td class="id">
      <a href="#" data-query="task" data-value="{id}">{id}</a>
      <a href="#" data-query="task-output-canvas" data-value="{id}">...</a>
    </td>
    <td class="started_at">{started_at}</td>
    <td class="last_activity_at">{last_activity_at}</td>
    <td class="user">
      <a href="#" data-query="event-user" data-value="{user}">{user}</a>
    </td>
    <td class="deployment">
      <a href="#" data-query="deployment" data-value="{deployment}">{deployment}</a>
      <a href="#" data-query="instances-canvas" data-value="{deployment}">...</a>
    </td>
    <td class="description">{description}</td>
    <td class="result">{result}</td>
  </tr>
</table>

<script type="text/javascript">

function TasksTable($el) {
  var dataSource = null;

  function setUp() {
    var moreCallback = function() { dataSource.More(); }
    var tmpls = {
      empty: Tmpl('<tr><td colspan="7">No matching tasks</td></tr>', []),
      error: Tmpl('<tr><td colspan="7">Error fetching tasks</td></tr>', []),
      dataItem: Tmpl($("#task-tmpl").html(),
        ["deployment", "description", "id",
          "last_activity_at", "result", "started_at", "state", "user"]),
    };
    var table = Table($el, moreCallback, tmpls);
    dataSource = TableDataSource("tasks", table, null);
  }

  setUp();

  return {
    Load: function() { dataSource.Load({"recent": "200"}); } // todo hard coded
  };
}

</script>
`
