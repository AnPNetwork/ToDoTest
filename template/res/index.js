class Task {

  constructor(id, description, status) {
    this.id = id;
    this.description = description;
    this.status = status;
  }

}

const ajax_url = './api';
var todolist = [];

function Trim(str) {
  return $.trim(str.replace(/\s+/g,' '));
}

$( document ).ready(function() {

  $('.ui.button.add').on("click", function() {
    $.modal('prompt',{
      title: 'Добавить задачу',
      placeholder: 'Введите описание задачи',
      handler: function(description){
          if (description == null) {
            return
          }

          ajaxAddTODO(description);
      }
    });
  });

  ajaxGetListTODO();
});

function renderTODO(data) {
  console.log(data);

  const result = JSON.parse(data);
  let tbody = $('table tbody');

  todolist = [];

  result.forEach(function(value) {
    todolist.push(new Task(value.id, value.description, value.status));
  });

  $(tbody).empty();

  todolist.forEach(function(value) {

    let element = `  
      <tr todo-id="${value.id}">
        <td class="description">       
          ${value.description}
        </td>
        <td class="status ${value.status ? 'set' : ''}">     
            ${value.status ? 'Да' : 'Нет'}     
        </td>
        <td class="delete">
          <i class="trash alternate icon" style="visibility: visible;"></i>
        </td>
      </tr>
    `;

    $(tbody).append(element);
  });

  $('table tr[todo-id] td.description').on("click", function() {
    let 
      tr = $(this).closest('tr[todo-id]'),
      id = parseInt($(tr).attr('todo-id')),
      description = Trim($(this).text()),
      status = $(tr).find('td.status').hasClass('set');

    $.modal('prompt',{
      title: 'Изменить описание задачи',
      placeholder: 'Введите описание задачи',
      defaultValue: description,
      handler: function(value){
  
        if (value == null) {
          return
        }

        description = value;

        ajaxChangeTODO(id, description, status);
      }
    });
  });

  $('table tr[todo-id] td.status').on("click", function() {
    let 
      tr = $(this).closest('tr[todo-id]'),
      id = parseInt($(tr).attr('todo-id')),
      description = Trim($(tr).find('.description').text()),
      status = $(this).hasClass('set');

    ajaxChangeTODO(id, description, !status);
  });

  $('table tr[todo-id] td.delete i').on("click", function() {
    let 
      tr = $(this).closest('tr[todo-id]'),
      id = parseInt($(tr).attr('todo-id')),
      description = Trim($(tr).find('.description').text());

    $.modal('confirm',{
      title: `Удалить задачу ${description} ?`,
      handler: function(choice){
        if (choice == 'Declined') {
          return
        }

        ajaxDeleteTODO(id, description)
      }
    });
  });
}

// ------------ AJAX ------------

function ajaxGetListTODO() {
  $.ajax({
    method: 'GET',
    url: ajax_url,
  })
  .done(function(resp) {
    renderTODO(resp);
  })
  .fail(function() {
    $.toast({
      class: 'error',
      message: `Ошибка получения списка задач - ${e.responseText}`
    });
  });
}

function ajaxAddTODO(description) {
  $.ajax({
    method: 'POST',
    url: ajax_url,
    data: JSON.stringify({description : description}),
    contentType: "application/json; charset=utf-8",
    dataType: "json",
    complete: function(resp) {
      console.log(resp);
  
      if (resp.status >= 200 && resp.status < 300) {
        $.toast({
          class: 'success',
          message: `Задача ${description} добавлена успешно`
        });
      } else {
        $.toast({
          class: 'error',
          message: `Ошибка добавления задачи - ${resp.responseText}`
        });
      }
  
      ajaxGetListTODO();
    },
  });
}

function ajaxChangeTODO(id, description, status) {

  $.ajax({
    method: 'PUT',
    url: ajax_url,
    data: JSON.stringify({
      id : id, 
      description: description,
      status: status,
    }),
    contentType: "application/json; charset=utf-8",
    dataType: "json",
    complete: function(resp) {
      console.log(resp);
  
      if (resp.status >= 200 && resp.status < 300) {
        $.toast({
          class: 'success',
          message: `Статус задачи ${description} изменен успешно`
        });
      } else {
        $.toast({
          class: 'error',
          message: `Ошибка изменения статуса задачи - ${resp.responseText}`
        });
      }
  
      ajaxGetListTODO();
    },
  });
}

function ajaxDeleteTODO(id, description) {
  $.ajax({
    method: 'DELETE',
    url: ajax_url,
    data: JSON.stringify({id : id}),
    contentType: "application/json; charset=utf-8",
    dataType: "json",
    complete: function(resp) {
      console.log(resp);
  
      if (resp.status >= 200 && resp.status < 300) {
        $.toast({
          class: 'success',
          message: `Задача ${description} добавлена успешно`
        });
      } else {
        $.toast({
          class: 'error',
          message: `Ошибка удаление задачи - ${resp.responseText}`
        });
      }
  
      ajaxGetListTODO();
    },
  });
}