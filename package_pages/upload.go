package package_pages

const UploadPage = `
<html>
<head>
  <style>
    label {
      display: inline-block;
      width: 55px;
    }

    #output {
      height: 500px;
      overflow-y: scroll;
    }
  </style>

  <title></title>
</head>
<body>
<form id="form" enctype="multipart/form-data" action="/upload" method="post" onsubmit="return send(event)">
  <label>file:</label> <input id="file" type="file" name="uploadFile" /><br/>
  <label>filepath:</label> <input id="filepath" name="filepath" title="filepath" placeholder="filepath" value="/tmp/upload.bin"><br/>
  <input id="formSubmit" type="submit" value="upload" />
</form>
<div id="output"></div>

<script>
  let form = document.getElementById('form');
  let output = document.getElementById('output');
  let filepath = document.getElementById('filepath');

  filepath.focus();
  filepath.select();

  function send(event) {
    output.innerText = "";
    event.preventDefault();
    let xhr = new XMLHttpRequest();

    xhr.open('POST', '/upload');

    let formData = new FormData(form);

    xhr.onprogress = function onprogress(e) {
      //var percentComplete = (e.position / e.totalSize)*100;
    };

    xhr.onload = function onload(e) {
      //
    };

    xhr.onreadystatechange = function onreadystatechange() {
      if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
        output.innerText = xhr.responseText;
      }
    };

    xhr.onerror = function onerror(e) {
      output.innerText = ("Error " + e.target.status + " occurred while receiving the document.");
    };

    xhr.send(formData);
    return false;
  }
</script>

</body>
</html>
`