active_sidebar("#qrcode", "#decoder");
$(function() {
    Dropzone.options.qrcodeDropzone = {
        url: "/qrcode/decode",
        method: "post",
        paramName: "qrcode",
        autoProcessQueue: false,
        uploadMultiple: false,
        parallelUploads: 1,
        maxFiles: 1,
        acceptedFiles: "image/*",
        maxFilesize: 5, // MB
        // Dropzone settings
        init: function() {
            var myDropzone = this;

            this.element.querySelector("button[type=submit]").addEventListener("click", function(e) {
                e.preventDefault();
                e.stopPropagation();
                myDropzone.processQueue();
            });
            this.element.querySelector("button[type=reset]").addEventListener("click", function(e) {
                myDropzone.removeAllFiles();
            });
            this.on("sendingmultiple", function() {
            });
            this.on("successmultiple", function(files, response) {
            });
            this.on("errormultiple", function(files, response) {
            });
            this.on("success", function(file, response) {
                console.log(typeof response);
                var title = "二维码解码器", message, method;
                if (response.status == "success") {
                    method = "success";
                    message = "二维码解码成功";
                    $("#content").val(response.content);
                } else {
                    method = "error";
                    message = "二维码解码失败: " + response.error;
                    $("#content").val("");
                }
                toastr.options = {
                    closeButton: true,
                    progressBar: true,
                    showMethod: 'slideDown',
                    timeOut: 3000
                };
                toastr[method](message, title);
                this.removeAllFiles();
            })
        }
    };
});