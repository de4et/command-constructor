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

function addOpenableBlockListeners(block, button, content) {
  const dropdown = document.querySelector(block);
  const dropdownButton = document.querySelector(button);
  const dropdownContent = document.querySelector(content);

  dropdownButton.addEventListener("click", (event) => {
    dropdownContent.classList.toggle("show");
  });

  document.addEventListener("click", (event) => {
    if (!dropdown.contains(event.target)) {
      dropdownContent.classList.remove("show");
    }
  });
}

function setError(parent, error_message_block, ...error_messages) {
  let error_block = error_message_block + "-list";

  let main_error_el = document.querySelector("." + error_block);
  console.log(main_error_el);

  if (main_error_el == null) {
    main_error_el = document.createElement("ul");
    main_error_el.classList.add(error_block);
    parent.appendChild(main_error_el);
  } else {
    console.log("hjer");
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

        console.log(Name + " : " + Value);
      }

      if (!formFilled) {
        setError(
          logForm,
          "log-error-message",
          "Все поля должны быть заполнены"
        );
        return;
      }

      // logForm.form.forEach((formEl) => {
      //   console.log(formEl);
      // });

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

// data = await getData()
