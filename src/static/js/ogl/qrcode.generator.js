active_sidebar("#qrcode", "#generator");
$(function() {
    $("#contentForm").validate({
        rules: {
            content: {
                required: true,
                maxlength: 250
            }
        },
        messages: {
            content: {
                required: "请输入文字内容",
                maxlength: "文字内容最多250个字符"
            }
        }
    });
    
    $("#confirm").click(function() {
        if ($("#contentForm").valid()) {
            $.ajax({
                type: "POST",
                url: "/qrcode/encode",
                data: $("#contentForm").serialize(),
                dataType: "json",
                success: function (data, textStatus, jqXHR) {
                    var status = data.status;
                    var title = "二维码生成器", message, method;
                    if (status == "success") {
                        method = "success";
                        message = "二维码生成成功";
                        $("#qrcode_img").attr("src", "data:image/png;base64," + data.base64);
                    } else {
                        method = "error";
                        message = "二维码生成失败: " + data.error;
                    }
                    toastr.options = {
                        closeButton: true,
                        progressBar: true,
                        showMethod: 'slideDown',
                        timeOut: 3000
                    };
                    toastr[method](message, title);
                },
                error: function (jqXHR, textStatus, errorThrown) {
                    console.log(textStatus);
                    console.log(errorThrown);
                }
            });
        }
    });
});