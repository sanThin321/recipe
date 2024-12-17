document.addEventListener("DOMContentLoaded", function() {
  window.onload
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
.then(res=> {
  console.log("Response status:", res.status);
  if (!res.ok) {
    return res.text().then(text => {
      throw new Error(`HTTP error! status: ${res.status}, res: ${text}`);
    });
  }
  return res.json(); // Expect JSON res
})
.then(data => {
  console.log("Server res:", data);
  if (data.status === "Recipe added") {
    alert("Recipe added successfully!");
    document.getElementById("addRecipeForm").reset();

  } else {
    alert("Failed to add recipe: " + (data.error || "Unknown error"));
  }
})
.catch(error => {
  console.error("Error adding recipe:", error);
  alert(`Error adding recipe: ${error.message}`);
});

}

