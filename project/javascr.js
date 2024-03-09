const tourContainer = document.getElementById("tourContainer");
const tourTitle = document.getElementById("tourTitle");
const tourDescription = document.getElementById("tourDescription");
const steps = tour.steps;
let currentStep = 0;

function showStep(step) {
  tourTitle.textContent = step.title;
  tourDescription.textContent = step.content;
}

function nextStep() {
  currentStep++;
  if (currentStep >= steps.length) currentStep = 0;
  showStep(steps[currentStep]);
}

function prevStep() {
  currentStep--;
  if (currentStep < 0) currentStep = steps.length - 1;
  showStep(steps[currentStep]);
}

document.getElementById("startTour").addEventListener("click", () => {
  tourContainer.style.display = "block";
  showStep(steps[currentStep]);
});

document.getElementById("next").addEventListener("click", nextStep);
document.getElementById("prev").addEventListener("click", prevStep);