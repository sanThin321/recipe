document.addEventListener("DOMContentLoaded", function() {
  const addRecipeForm = document.getElementById("addRecipeForm");
  addRecipeForm.addEventListener("submit", function(event) {
    event.preventDefault();
    addRecipe();
  });
});

function addRecipe() {
  const formData = new FormData(document.getElementById("addRecipeForm"));
  console.log("Submitting recipe...");
  
  fetch("/recipe", {
  method: "POST",
  body: formData,
})
.then(response => {
  console.log("Response status:", response.status);
  if (!response.ok) {
    return response.text().then(text => {
      throw new Error(`HTTP error! status: ${response.status}, response: ${text}`);
    });
  }
  return response.json(); // Expect JSON response
})
.then(data => {
  console.log("Server response:", data);
  if (data.status === "Recipe added") {
    alert("Recipe added successfully!");
  } else {
    alert("Failed to add recipe: " + (data.error || "Unknown error"));
  }
})
.catch(error => {
  console.error("Error adding recipe:", error);
  alert(`Error adding recipe: ${error.message}`);
});

}

function newRecipe(recipe) {
  const main = document.querySelector("main");
  
  const container = document.createElement("div");
  container.classList.add("card_container");
  main.appendChild(container);

  const recipeImage = document.createElement("div");
  recipeImage.classList.add("recipe_img");
  container.appendChild(recipeImage);

  const img = document.createElement("img");
  img.src = `data:image/jpeg;base64,${recipe.image}`;
  img.alt = "recipe image";
  recipeImage.appendChild(img);

  const cardBody = document.createElement("div");
  cardBody.classList.add("card_body");
  container.appendChild(cardBody);

  const recipeName = document.createElement("h3");
  recipeName.classList.add("recipe_name");
  recipeName.textContent = recipe.recipename;
  cardBody.appendChild(recipeName);

  const viewMoreBtn = document.createElement("button");
  viewMoreBtn.textContent = "View More";
  cardBody.appendChild(viewMoreBtn);
}
