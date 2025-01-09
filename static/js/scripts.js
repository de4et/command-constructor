async function getData(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const json = await response.json();
    console.log(json);
  } catch (error) {
    console.error(error.message);
  }
}

function addFocusListeners(button, block) {
  const dropdown = document.querySelector(block);
  const dropdownButton = document.querySelector(button);

  dropdownButton.addEventListener("click", (event) => {
    dropdown.classList.toggle("focus");
  });

  document.addEventListener("click", (event) => {
    if (!dropdownButton.contains(event.target)) {
      dropdown.classList.remove("focus");
    }
  });
}

openableBlocks = {};

function addOpenableBlockListeners(block, button, content) {
  const dropdown = document.querySelector(block);
  const dropdownButton = document.querySelector(button);
  const dropdownContent = document.querySelector(content);

  dropdownButton.addEventListener("click", (event) => {
    dropdownContent.classList.toggle("show");
  });

  if (openableBlocks[block] == undefined) {
    openableBlocks[block] = [];
  }
  openableBlocks[block].push(dropdownButton);

  document.addEventListener("click", (event) => {
    if (
      !dropdown.contains(event.target) &&
      !openableBlocks[block].includes(event.target)
    ) {
      dropdownContent.classList.remove("show");
    }
  });
}

function setError(parent, error_message_block, ...error_messages) {
  let error_block = error_message_block + "-list";

  let main_error_el = document.querySelector("." + error_block);

  if (main_error_el == null) {
    main_error_el = document.createElement("ul");
    main_error_el.classList.add(error_block);
    parent.appendChild(main_error_el);
  } else {
    main_error_el.innerHTML = "";
  }

  for (var i = 0; i < error_messages.length; i++) {
    if (error_messages[i] == null) continue;
    var element = document.createElement("li");
    element.classList.add(error_message_block);
    element.appendChild(document.createTextNode(error_messages[i]));
    main_error_el.appendChild(element);
  }
}

document.addEventListener("DOMContentLoaded", (event) => {
  document.querySelectorAll(".login_form").forEach((logForm) => {
    logForm.addEventListener("submit", (event) => {
      event.preventDefault();
      const formData = new FormData(logForm);

      let formFilled = true;
      for (I = 0; I < logForm.length; I++) {
        var Name = logForm[I].getAttribute("name");
        var Value = logForm[I].value;

        if (Name == null) {
          continue;
        }

        if (Value == "") {
          formFilled = false;
          break;
        }
      }

      if (!formFilled) {
        setError(
          logForm,
          "log-error-message",
          "Все поля должны быть заполнены"
        );
        return;
      }

      const button = logForm.querySelector("button");
      const loading_animation = logForm.querySelector(".loading-animation");

      button.hidden = true;
      loading_animation.hidden = false;

      const url = logForm.getAttribute("action");
      fetch(url, {
        method: "POST",
        body: formData,
      }).then((res) => {
        let status = res.status;
        res.json().then((data) => {
          if (status != 200) {
            button.hidden = false;
            loading_animation.hidden = true;
            if (data.error) {
              setError(logForm, "log-error-message", data.error);
            } else {
              setError(
                logForm,
                "log-error-message",
                data.name,
                data.email,
                data.password
              );
            }
          } else {
            window.location.reload();
          }
        });
      });
    });
  });
});

function selectArgumentTypeChange(event) {
  const selectValue = event.target.value; // id of template
  var template = document.getElementById(selectValue);
  const item = template.content.cloneNode(true);
  var argument_block = event.target.closest(".argument-edit");
  lastChild = argument_block.querySelector(".argument-type");

  while (lastChild.nextSibling) {
    argument_block.removeChild(lastChild.nextSibling);
  }
  lastChild.after(item);
  updatePreviewCommand();
}

function addArgumentClick(event) {
  event.preventDefault();

  var argument_id = generateArgumentID();

  addNewArgumentBlock(argument_id);
  updatePreviewCommand();
}

function fillCommandTemplate(commandTemplate) {
  // fill main inputs
  document.querySelector(".form-input-name").value =
    commandTemplate.get("name");
  document.querySelector(".form-input-description").value =
    commandTemplate.get("description");
  document.querySelector(".form-input-commandName").value =
    commandTemplate.get("commandName");

  // fill arguments
  commandTemplate.get("templateParams").forEach((param) => {
    var argument_id = generateArgumentID();
    argEl = addNewArgumentBlock(argument_id);
    fillArgument(argEl, param);
  });

  updatePreviewCommand();
}

function fillArgument(elem, arg) {
  select = elem.querySelector("select");
  $(select).val(
    select.querySelector(
      "option[data-type-id='" + arg.get("type").toString() + "']"
    ).value
  );
  $(select).change();

  nameEl = elem.querySelector(".argument-name-input");
  if (nameEl) nameEl.value = arg.get("name");

  description = elem.querySelector(".argument-description-textarea");
  description.value = arg.get("description");

  isconstant = elem.querySelector(".argument-isconstant-checkbox");
  console.log(isconstant);
  isconstant.checked = arg.get("isconstant");

  default_value = elem.querySelector(".argument-default-value");
  if (default_value) default_value.value = arg.get("defaultValue");

  buttonEl = elem.querySelector(".dropdown-add-value-button");
  if (arg.get("type") == "3") {
    arg.get("value").forEach((val) => addDropdownValue(buttonEl, val));
  }
}

var lastArgumentID = 0;

function addNewArgumentBlock(argument_id) {
  var arguments_block = document.querySelector(".create-form-arguments");
  var button = arguments_block.querySelector(".add-argument-button");

  var template = document.querySelector("#argument-template");
  var item = template.content.cloneNode(true);
  argBlock = item.querySelector("div");
  item.querySelector("div").id = "argument-" + argument_id;
  button.before(item);
  return argBlock;
}

function generateArgumentID() {
  return lastArgumentID++;
}

function updatePreviewCommand() {
  var previewEl = document.querySelector(".preview-command");

  const commandName = document.getElementsByName("commandName")[0].value;
  const argumentsBlocks = document
    .querySelector(".create-form-arguments")
    .querySelectorAll(".argument-edit");

  previewCommandNameEl = previewEl.getElementsByClassName(
    "preview-command-commandname"
  )[0];
  previewCommandNameEl.textContent = commandName;
  if (previewCommandNameEl.textContent == "") {
    previewCommandNameEl.textContent = "имя";
  }

  while (previewCommandNameEl.nextSibling) {
    if (previewCommandNameEl.nextSibling.tagName == "TEMPLATE") break;
    previewEl.removeChild(previewCommandNameEl.nextSibling);
  }

  args = argumentsToMap();
  args.forEach((arg) => {
    elem = getPreviewParamElem(arg);
    previewEl.appendChild(elem);
  });
}

function getPreviewParamElem(arg) {
  var paramTemplate = document.querySelector("#preview-param-template");
  var itemParam = paramTemplate.content.cloneNode(true).firstElementChild;

  descriptionEl = itemParam.querySelector(".preview-param-description");
  descriptionEl.textContent = arg.get("description");

  if (arg.get("isconstant")) {
    itemParam.setAttribute("data-paramtype", "constant");
    itemParam.classList.add("constant");

    sParamTemplate = document.querySelector("#preview-param-constant-template");
    itemSParam = sParamTemplate.content.cloneNode(true).firstElementChild;
    itemSParam.textContent = [arg.get("name"), arg.get("defaultValue")].join(
      " "
    );
    itemParam.appendChild(itemSParam);
    return itemParam;
  }

  itemParam.setAttribute("data-paramtype", "template");
  switch (arg.get("type")) {
    case 0:
      sParamTemplate = document.querySelector("#preview-param-string-template");
      itemSParam = sParamTemplate.content.cloneNode(true);
      itemSParam.querySelector("span").textContent = arg.get("name");
      if (itemSParam.querySelector("span").textContent == "") {
        itemSParam.querySelector("span").remove();
      }
      itemSParam.querySelector("input").value = arg.get("defaultValue");
      break;
    case 1:
      sParamTemplate = document.querySelector("#preview-param-empty-template");
      itemSParam = sParamTemplate.content.cloneNode(true);
      itemSParam.querySelector("span").textContent = arg.get("name");
      if (itemSParam.querySelector("span").textContent == "") {
        itemSParam.querySelector("span").remove();
      }

      break;
    case 2:
      sParamTemplate = document.querySelector(
        "#preview-param-nameless-template"
      );
      itemSParam = sParamTemplate.content.cloneNode(true);
      itemSParam.querySelector("input").value = arg.get("defaultValue");
      break;
    case 3:
      sParamTemplate = document.querySelector("#preview-param-popup-template");
      itemSParam = sParamTemplate.content.cloneNode(true);
      itemSParam.querySelector("span").textContent = arg.get("name");
      if (itemSParam.querySelector("span").textContent == "") {
        itemSParam.querySelector("span").remove();
      }
      fillSelectWithValues(
        itemSParam.querySelector(".dropdown-menu"),
        arg.get("value")
      );
      break;
  }
  itemParam.prepend(itemSParam);

  return itemParam;
}

function fillSelectWithValues(selectEl, values) {
  var button = selectEl.closest(".dropdown").querySelector("button");
  if (values != 0) {
    button.textContent = values[0];
  } else {
    button.textContent = "";
  }

  for (var i = 0; i < values.length; i++) {
    var val = values[i];

    var paramTemplate = document.querySelector("#param-value-template");
    var itemParam = paramTemplate.content.cloneNode(true).firstElementChild;
    var el = itemParam.querySelector("a");
    el.text = val;

    $(el).click(function () {
      var selText = $(this).text();
      $(this).parents(".dropdown").find(".dropdown-toggle").text(selText);
    });

    selectEl.appendChild(itemParam);
  }
}

function commandNameInput(event) {
  updatePreviewCommand();
}

function argumentsToMap() {
  var args = [];

  Array.from(document.getElementsByClassName("argument-edit")).forEach(
    (elem) => {
      args.push(argumentToMap(elem));
    }
  );
  return args;
}

function argumentToMap(elem) {
  typeEl = elem.querySelector(".select-argument-type");
  var type = typeEl.options[typeEl.selectedIndex].getAttribute("data-type-id");

  default_value =
    elem.querySelector(".argument-default-value") != null
      ? elem.querySelector(".argument-default-value").value
      : "";
  isconstant = elem.querySelector(".argument-isconstant-checkbox").checked;
  description = elem.querySelector(".argument-description-textarea").value;
  name =
    elem.querySelector(".argument-name-input") != null
      ? elem.querySelector(".argument-name-input").value
      : "";

  value = [];
  if (type == "3") {
    elem.querySelectorAll(".dropdown-value").forEach((valEl) => {
      value.push(valEl.querySelector("input").value);
    });
  }

  m = new Map();
  m.set("name", name);
  m.set("description", description);
  m.set("type", parseInt(type));
  m.set("defaultValue", default_value);
  m.set("value", value);
  m.set("isconstant", isconstant);
  return m;
}

function addDropdownValueClick(e) {
  e.preventDefault();
  var buttonEl = e.currentTarget;

  var template = document.querySelector("#argument-dropdown-value-template");
  var item = template.content.cloneNode(true).firstElementChild;

  buttonEl.before(item);
  updatePreviewCommand();
}

function addDropdownValue(buttonEl, value) {
  var template = document.querySelector("#argument-dropdown-value-template");
  var item = template.content.cloneNode(true).firstElementChild;
  item.querySelector("input").value = value;

  buttonEl.before(item);
}

function deleteDropdownValue(event) {
  event.preventDefault();
  var buttonEl = event.target;
  var valueEl = buttonEl.closest(".dropdown-value");
  valueEl.remove();
  updatePreviewCommand();
}

function deleteArgumentClick(event) {
  event.preventDefault();
  const buttonEl = event.target;
  argumentBlock = buttonEl.closest(".argument-edit");
  argumentBlock.remove();
  updatePreviewCommand();
}

let newX,
  newY,
  startX,
  startY = 0;
let relativePos = {};
let initialPos = {};

function argumentMouseDown(event) {
  const buttonEl = event.currentTarget;
  argumentBlock = buttonEl.closest(".argument-edit");

  const parentPos = argumentBlock.getBoundingClientRect();
  const childPos = buttonEl.getBoundingClientRect();
  initialPos = parentPos;

  relativePos.top = childPos.top - parentPos.top + childPos.height / 2;
  relativePos.left = childPos.left - parentPos.left + childPos.width / 2;

  argumentBlock.classList.add("dragging");

  startX = event.clientX;
  startY = event.clientY;

  document.addEventListener("mousemove", argumentMouseMove);
  document.addEventListener("mouseup", argumentMouseUp);
  argumentMouseMove(event);

  // add targets
  var template = document.querySelector("#argument-target-template");
  document.querySelectorAll(".argument-edit").forEach((elem) => {
    if (elem.classList.contains("dragging")) return;
    var item = template.content.cloneNode(true).firstElementChild;
    elem.after(item);
  });

  var item = template.content.cloneNode(true).firstElementChild;
  argsBlock = document.querySelector(".create-form-arguments");
  argsBlock.prepend(item);
}

function argumentMouseMove(event) {
  newX = startX - event.clientX;
  newY = startY - event.clientY;

  startX = event.clientX;
  startY = event.clientY;

  upPos = {};
  upPos.x = event.clientX;
  upPos.y = event.clientY;

  const isInside = (point, rect) =>
    point.x > rect.left &&
    point.x < rect.right &&
    point.y > rect.top &&
    point.y < rect.bottom;

  document.querySelectorAll(".argument-target").forEach((elem) => {
    var elRect = elem.getBoundingClientRect();
    elSvg = elem.querySelector("svg");
    elSvg.style.fill = "white";
    if (isInside(upPos, elRect)) {
      elSvg.style.fill = "orange";
    }
  });

  argumentBlock = document.querySelector(".dragging");
  argumentBlock.style.top = startY - relativePos.top + "px";
  argumentBlock.style.left = startX - relativePos.left + "px";
}

function argumentMouseUp(event) {
  document.removeEventListener("mousemove", argumentMouseMove);
  document.removeEventListener("mouseup", argumentMouseUp);

  upPos = {};
  upPos.x = event.clientX;
  upPos.y = event.clientY;

  const isInside = (point, rect) =>
    point.x > rect.left &&
    point.x < rect.right &&
    point.y > rect.top &&
    point.y < rect.bottom;

  document.querySelectorAll(".argument-target").forEach((elem) => {
    var elRect = elem.getBoundingClientRect();
    if (isInside(upPos, elRect)) {
      elem.after(argumentBlock);
    }
  });

  document.querySelectorAll(".argument-target").forEach((elem) => {
    elem.remove();
  });

  argumentBlock.style = "";
  argumentBlock.classList.remove("dragging");

  updatePreviewCommand();
}

function makeUnselectable(elem) {
  if (typeof elem == "string") elem = document.getElementById(elem);
  if (elem) {
    elem.onselectstart = function () {
      return false;
    };
    elem.style.MozUserSelect = "none";
    elem.style.KhtmlUserSelect = "none";
    elem.unselectable = "on";
  }
}

function excludeParamClick(event) {
  elem = event.currentTarget;
  if (event.shiftKey) {
    event.preventDefault();
    event.stopPropagation();
    event.stopImmediatePropagation();
    elem.classList.toggle("param-excluded");

    childInput = elem.querySelector("input");
    if (childInput) childInput.toggleAttribute("disabled");

    childDropdown = elem.querySelector(".dropdown-toggle");
    if (childDropdown) {
      childDropdown.classList.toggle("disabled");
      disabled = childDropdown.classList.contains("disabled");
      $(childDropdown).dropdown("toggle");
      childDropdown.addEventListener(
        "show.bs.dropdown",
        (event) => {
          if (disabled) {
            event.preventDefault();
          }
        },
        { once: true }
      );
    }
  }
}
function toMap(obj) {
  if (Array.isArray(obj)) {
    // If it's an array, recursively convert each element
    return obj.map(toMap);
  } else if (obj !== null && typeof obj === "object") {
    // If it's an object, convert it into a Map
    const map = new Map();
    for (const [key, value] of Object.entries(obj)) {
      map.set(key, toMap(value)); // Recursively process values
    }
    return map;
  }
  // If it's neither an array nor an object, return it as-is (primitive values)
  return obj;
}

function createTemplateClick(event) {
  form = document.querySelector(".create-form");
  formdata = new FormData(form);
  mapdata = new Map(formdata);
  args = argumentsToMap();
  mapdata.set("templateParams", args);

  const serializedMap = Object.fromEntries(
    Array.from(mapdata.entries()).map(([key, value]) => {
      if (Array.isArray(value)) {
        return [key, value.map((innerMap) => Object.fromEntries(innerMap))];
      }
      return [key, value];
    })
  );

  const button = form.querySelector(".create-form-button");
  const loading_animation = form.querySelector(".loading-animation");

  button.hidden = true;
  loading_animation.hidden = false;

  const url = form.getAttribute("action");
  fetch(url, {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(serializedMap),
  }).then((res) => {
    let status = res.status;
    res.json().then((data) => {
      button.hidden = false;
      loading_animation.hidden = true;
      if (status != 200) {
        setError(form, "log-error-message", data.error);
      } else {
        window.location.href = window.location.origin + "/edit/" + data.id;
      }
    });
  });
}

function editTemplateClick(event) {
  form = document.querySelector(".create-form");
  formdata = new FormData(form);
  mapdata = new Map(formdata);
  args = argumentsToMap();
  mapdata.set("templateParams", args);

  const serializedMap = Object.fromEntries(
    Array.from(mapdata.entries()).map(([key, value]) => {
      if (Array.isArray(value)) {
        return [key, value.map((innerMap) => Object.fromEntries(innerMap))];
      }
      return [key, value];
    })
  );

  const queryString = window.location.href;
  const segments = new URL(queryString).pathname.split("/");
  const id = segments.pop() || segments.pop();

  const button = form.querySelector(".create-form-button");
  const loading_animation = form.querySelector(".loading-animation");

  button.hidden = true;
  loading_animation.hidden = false;

  const url =
    window.location.origin + "/" + form.getAttribute("action") + "/" + id;
  fetch(url, {
    method: "PUT",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(serializedMap),
  }).then((res) => {
    let status = res.status;
    res.json().then((data) => {
      button.hidden = false;
      loading_animation.hidden = true;
      if (status != 200) {
        setError(form, "log-error-message", data.error);
      } else {
        window.location.reload();
      }
    });
  });
}

$(document).ready(function () {
  $(".create-form").bind("keypress", function (e) {
    if (e.keyCode == 13 && e.target.tagName != "TEXTAREA") {
      return false;
    }
  });
});
