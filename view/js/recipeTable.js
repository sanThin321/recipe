// Ensure that the DOMContentLoaded event listener is properly set up
document.addEventListener("DOMContentLoaded", function() {
    const addRecipeForm = document.getElementById("addRecipeForm");
    addRecipeForm.addEventListener("submit", function(event) {
      event.preventDefault();
      addRecipe();
    });
  });
  
  // Call the fetchRecipes function when the window loads
  window.onload = fetchRecipes;
  
  // Function to fetch recipes from the server
  function fetchRecipes() {
    fetch("/recipes")
        .then((response) => response.json()) // Parse the response as JSON
        .then((data) => showRecipes(data))
        .catch((error) => console.error('Error fetching recipes:', error));
  }
  
  // Function to display recipes in the table
  function showRecipes(data) {
    const table = document.getElementById("myTable");
    
    // Clear existing table rows
    table.innerHTML = "<tr><th>RID</th><th>Recipe Name</th><th>Action</th></tr>";
  
    data.forEach(recipe => {
      addRecipeTable(recipe);
    });
  }
  
  // Function to add a recipe to the table
  function addRecipeTable(recipe) {
    const table = document.getElementById("myTable");
  
    // Create a new row for the recipe
    const row = document.createElement("tr");
  
    // Add cells for recipe ID, name, and action
    const ridCell = document.createElement("td");
    ridCell.textContent = recipe.rid;
    row.appendChild(ridCell);
  
    const nameCell = document.createElement("td");
    nameCell.textContent = recipe.recipename;
    row.appendChild(nameCell);
  
    const actionCell = document.createElement("td");
    const editButton = document.createElement("a");
    editButton.href = `recipe_edit.html?id=${recipe.rid}`;
    editButton.textContent = "Edit";
    editButton.classList.add("btn");
    actionCell.appendChild(editButton);
    row.appendChild(actionCell);
  
    // Append the row to the table
    table.appendChild(row);
  }
  
  // Other functions remain unchanged...
  