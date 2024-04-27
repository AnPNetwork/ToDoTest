$( document ).ready(function() {
   let container_todo =  $("#container_todo");
   $(container_todo).empty();

   GetTODO()

 });


function GetTODO() {
  var jqxhr = $.ajax( "example.php" )
  .done(function(data) {
    console.log(data);
  })
  .fail(function() {
    alert( "error" );
  })
  .always(function() {
    alert( "complete" );
  });
 }