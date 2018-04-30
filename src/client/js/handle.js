    function dragenter(e) {
        e.stopPropagation();
        e.preventDefault();
    }

    function dragover(e) {
        e.stopPropagation();
        e.preventDefault();
    }

    function drop(e) {
        e.stopPropagation();
        e.preventDefault();
        var dt = e.dataTransfer;
        files = dt.files;
        //get dropbox
        var dropbox = document.getElementById('dropzone');
        dropbox.innerHTML = '';
        dropbox.setAttribute('style', "line-height: 10px");
        // show file list on dropbox
        for (var i = 0; i < files.length; i++) {
            var p = document.createElement('p');
            p.innerHTML = files[i].name;
            p.setAttribute('class', 'filelist');
            dropbox.appendChild(p);
        }
    }

    //handle files
    function handleFile() {
        isExist();
    }

    // check file is exist
    function isExist() {
        var xhr = new XMLHttpRequest();
        var file_name_array = [];

        function disp_confirm(file_name) {
            var r = confirm(file_name + " already exist, do you want to overwrite?")
            if (r == false) {
                banFilesName.push(file_name);
            };
        }

        xhr.onload = function () {
            var data = JSON.parse(this.responseText);
            for (var i = 0; i < data.length; i++) {
                if (data[i].IsExist == true) {
                    //                    alert(data[i].Fn + ' already exist');
                    disp_confirm(data[i].Fn);
                }
            }
            uploadFile();
        }
        for (var i = 0; i < files.length; i++) {
            file_name_array.push(files[i].name);
        }

        var list_text = JSON.stringify(file_name_array);
        var obj = '{' + '"fns":' + list_text + '}'
        //ajax send request
        xhr.open('post', 'http://127.0.0.1:8080/api/isExist');
        xhr.send(obj);
    }

    // check elem is in list
    function isInList(list, elem) {
        for (var i = 0; i < list.length; i++) {
            if (elem == list[i]) {
                return true;
            }
        }
        return false;
    }

    //upload file ro server
    function uploadFile() {
        //initial progress value
        var dropbox = document.getElementById('dropzone');
        dropbox.innerHTML = '';
        dropbox.setAttribute('style', "line-height: 300px");
        //create progress
        var progress = document.createElement('progress');
        progress.setAttribute('min', '0');
        progress.setAttribute('max', '100');
        progress.setAttribute('value', "0");
        dropbox.appendChild(progress);
        progress.value = 0;

        var formData = new FormData();
        var xhr = new XMLHttpRequest();
        xhr.upload.onprogress = function (event) {　　
            if (event.lengthComputable) {　　　　　　
                var complete = (event.loaded / event.total * 100 | 0);　　　　　　　　　　　
                progress.value = progress.innerHTML = complete;　　
            }
        };
        xhr.onload = function () {
            var r = confirm("Upload Complete");
            if (r == true) {
                location.reload();
            }
        }
        for (var i = 0; i < files.length; i++) {
            if (isInList(banFilesName, files[i].name)) {
                continue;
            }
            formData.append("uploadfile", files[i]);
            console.log('file : ' + files[i].name);
        }
        //ajax send request
        xhr.open('post', 'http://127.0.0.1:8080/api/upload');
        xhr.send(formData);
        banFilesName = [];
    }

    var dropzone;
    var files;
    var banFilesName = [];
    dropzone = document.getElementById("dropzone");
    dropzone.addEventListener("dragenter", dragenter, false);
    dropzone.addEventListener("dragover", dragover, false);
    dropzone.addEventListener("drop", drop, false);
