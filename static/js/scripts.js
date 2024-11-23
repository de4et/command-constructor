function profile_click() {
  const target = document.getElementById("profile_menu");
  const button = document.getElementById("profile_button");
  if (target.style.display === "none") {
    target.style.display = "block";
    button.focus();
  } else {
    target.style.display = "none";
    button.blur();
  }
  // document.getElementById("profile_menu").style.display = "block";
  console.log("clicked");
}

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

function focusout_profile() {
  const target = document.getElementById("profile_menu");
  console.log("ima here");
  target.style.display = "none";
}
// data = await getData()

document.addEventListener("click", function (e) {
  const target = document.getElementById("profile_menu");
  if (
    !document.getElementById("profile_menu").contains(e.target) &&
    !document.getElementById("profile_button").contains(e.target)
  ) {
    target.style.display = "none";
  }
});
