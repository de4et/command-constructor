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

// data = await getData()
document.addEventListener("DOMContentLoaded", function (event) {
  const dropdown = document.querySelector(".profile_block");
  const dropdownButton = document.querySelector(".profile_button");
  const dropdownContent = document.querySelector(".profile_menu_list");

  // Обработчик для кнопки
  dropdownButton.addEventListener("click", (event) => {
    event.stopPropagation();
    dropdownContent.classList.toggle("show");
  });

  // Закрытие при клике вне dropdown
  document.addEventListener("click", () => {
    dropdownContent.classList.remove("show");
  });
});
