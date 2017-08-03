package package_pages

const RunPage = `
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
<form id="form" action="/run" onsubmit="return send(event)">
    <label>cmd:</label> <input id="cmd" name="cmd" title="cmd" placeholder="cmd"><br/>
    <br/>
    <label>timeout:</label> <input id="timeout" name="timeout" title="timeout" placeholder="timeout" type="number"
                                   value="30000"><br/>
    <label>nowait:</label> <input id="nowait" name="nowait" title="nowait" placeholder="nowait" type="checkbox"><br/>
    <br/>
    <label>params:</label> <input id="params0" name="params[]" title="params" placeholder="params"><br/>
    <label>params:</label> <input id="params1" name="params[]" title="params" placeholder="params"><br/>
    <input id="formSubmit" type="submit" value="submit">
</form>

<div id="output"></div>

<script>
  let form = document.getElementById('form');
  let output = document.getElementById('output');

  function formatParams() {
    let res = "?";
    for (let inp of form.querySelectorAll('input:not([type=submit])')) {
      let name = inp.name;//.replace("[]", "");
      let val = "wrong_value";
      switch (inp.type) {
        case "checkbox":
          val = inp.checked ? "1" : "";
          break;
        default:
          val = inp.value;
      }
      val = encodeURIComponent(val);
      res += (name + '=' + val + '&')
    }
    return res;
  }


  function send(event) {
    output.innerText = "";
    event.preventDefault();
    let xhr = new XMLHttpRequest();

    xhr.open('GET', '/run'/* + formatParams()*/);

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

    xhr.send(/*formData*/null);
    return false;
  }
</script>

</body>
</html>
`