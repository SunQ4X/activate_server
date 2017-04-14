

function login() 
{	
	var req = {
		username: $("#usr").val(),
		password: $("#pwd").val()
	};
	
	$.ajax({
		type: "POST",
		url: "/login",
		data: JSON.stringify(req),
		dataType: "json",
		processData:false,
		success: function(data){
			if (0 == data.ErrorCode){
						
			} else {
				alert(data.ErrorCode)
			}
		}
	});
}