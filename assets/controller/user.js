loadAllUsers();

//Save User
$("#btnAddUser").click(function () {
    let id = $("#id").val();
    let username = $("#username").val();
    let email = $("#email").val();
    let address = $("#address").val();

    $.ajax({
        method: "post",
        url: "http://localhost:8080/api/users",
        contentType: "application/json",
        data: JSON.stringify({
            "id": id,
            "username": username,
            "email": email,
            "address": address
        }),

        success: function (res) {
            alert("User Registered");
            loadAllUsers()
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });

});


//update User
$("#btnUpdateUser").click(function () {
    let id = $("#id").val();
    let username = $("#username").val();
    let email = $("#email").val();
    let address = $("#address").val();

    $.ajax({
        method: "put",
        url: "http://localhost:8080/api/users/" + id,
        contentType: "application/json",
        data: JSON.stringify({
            "id": id,
            "username": username,
            "email": email,
            "address": address

        }),
        success: function (res) {
            alert("User Updated");
            loadAllUsers()
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });

});


//delete user
$("#btnDeleteUser").click(function (){
    let id=$("#id").val();
    $.ajax({
        method:"DELETE",
        url:"http://localhost:8080/api/users/"+ id,
        success:function (res){
            alert("the user is removed");
            loadAllUsers()
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });
});

function loadAllUsers() {
    $('#tblUser').empty();

    $.ajax({
        method: "GET",
        url: "http://localhost:8080/api/users",
        success: function (res) {
            let data = res.data;
            console.log(data)
            for (var i in res){
                let id = res[i].id;
                let username = res[i].username;
                let email = res[i].email;
                let address = res[i].address;

                var row=`<tr> <td>${id}</td> <td>${username}</td><td>${email}</td><td>${address}</td></tr>`;
                $('#tblUser').append(row);
            }
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    })
}

// search user
$("#btnSearchUser").click(function () {
    let id = $("#id").val();
    $.ajax({
        method: "GET",
        url: "http://localhost:8080/api/users/" + id,
        success: function (res) {
            console.log(res);
            let c = res.data;
            $("#username").val(res.username);
            $("#email").val(res.email);
            $("#address").val(res.address);
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });
});
