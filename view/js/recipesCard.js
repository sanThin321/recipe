document.addEventListener('DOMContentLoaded', function () {
  const recipeId = getRecipeIdFromURL();

  if (recipeId) {
      fetchRecipe(recipeId);
  }
})
window.onload = function() {
    fetch("/recipes")
        .then((response) => response.text()) // parse the JSON directly
        .then((data) => showRecipes(data))
        .catch((error) => console.error('Error fetching recipes:', error));
  }
  function showRecipes(data){
    // console.log("Data received:", data);
    const recipes=JSON.parse(data);
    // console.log("pass received:", data);
    // console.log(recipes)
    recipes.forEach(recipe=>{
      newRecipe(recipe);
    })
  }
  function getRecipeIdFromURL() {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get('id');
}
function fetchRecipe(recipeId) {
  console.log(recipeId)
  fetch(`/recipe/` + recipeId)
      .then(response => response.json())
      .then(data => {
          showRecipe(data);
      })
      .catch(error => {
          console.error('Error fetching recipe:', error);
      });
}
function showRecipe(recipe) {
  const pop = document.getElementById("pop");
  const popup = document.createElement("div");
  popup.classList.add("popup");
  pop.appendChild(popup);

  const contain = document.createElement("div");
  contain.classList.add("contain");
  popup.appendChild(contain);

  const ingredientsTopic = document.createElement("h4");
  ingredientsTopic.textContent = "Ingredients";
  ingredientsTopic.classList.add("ingredients_topic");
  contain.appendChild(ingredientsTopic);

  const ingredients = document.createElement("p");
  ingredients.textContent = recipe.ingredient;
  contain.appendChild(ingredients);

  const stepsTopic = document.createElement("h4");
  stepsTopic.textContent = "Steps";
  stepsTopic.classList.add("steps_topic");
  contain.appendChild(stepsTopic);

  const steps = document.createElement("p");
  steps.textContent = recipe.steps;
  contain.appendChild(steps);

  // Close button for the popup
  const closeBtn = document.createElement("button");
  closeBtn.textContent = "Close";
  closeBtn.addEventListener("click", () => {
    pop.classList.remove("show")
    pop.removeChild(popup)
  });
  contain.appendChild(closeBtn);
  console.log("hello")
  pop.classList.add("show");
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
  
    const card_detail=document.createElement('div')
    card_detail.classList.add("card-detail")
    container.appendChild(card_detail)
    const cardBody = document.createElement("div");
    cardBody.classList.add("card_body");
    card_detail.appendChild(cardBody);
  
    const recipeName = document.createElement("h3");
    recipeName.classList.add("recipe_name");
    recipeName.textContent = recipe.recipename;
    cardBody.appendChild(recipeName);
  
    const viewMoreBtn = document.createElement("button");
    viewMoreBtn.classList.add("view")
    viewMoreBtn.textContent = "View More";
    cardBody.appendChild(viewMoreBtn);

    viewMoreBtn.addEventListener("click", () => {
      fetchRecipe(recipe.rid)
  });
  }
  