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

document.addEventListener("DOMContentLoaded", (event) => {
  document.querySelectorAll(".login_form").forEach((logForm) => {
    logForm.addEventListener("submit", (event) => {
      event.preventDefault();

      const button = logForm.querySelector("button");
      const loading_animation = logForm.querySelector(".loading-animation");

      button.hidden = true;
      loading_animation.hidden = false;

      const formData = new FormData(logForm);

      const url = logForm.getAttribute("action");
      fetch(url, {
        method: "POST",
        body: formData,
      }).then((res) => {
        let status = res.status;
        res.json().then((data) => {
          console.log(status);
          console.log(data);
          if (status != 200) {
            button.hidden = false;
            loading_animation.hidden = true;
          } else {
            window.location.reload();
          }
        });
      });
    });
  });
});

// data = await getData()
