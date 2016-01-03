
// Global Variables
var canvas, context, mouseCords, clientX = 0, clientY = 0;

/******  Helper Functions ******/
function getMousePosition(canvs) {
	var rect = canvs.getBoundingClientRect();
	return {
		x: Math.round((clientX-rect.left)/(rect.right-rect.left)*canvs.width),
		y: Math.round((clientY-rect.top)/(rect.bottom-rect.top)*canvs.height)
	}
}

function canvasMouseEvent(evt) {
	clientX = evt.clientX;
	clientY = evt.clientY;
}

/******  UI Functions **********/

function animationObj() {
	this.x = 0,
	this.y = 0,
	this.z = 0,
	this.draw = null,
	this.update = null,
};

/******  Main Functions ********/

function init(canvsId) {
	canvas = document.getElementById(canvsId);
	context = canvas.getContext('2d');
	canvas.addEventListener('mousemove', canvasMouseEvent);
	MainLoop.setUpdate(updateFrame).setDraw(drawFrame).setBegin(beginFrame).start();
}

function beginFrame() {
	mouseCords = getMousePosition(canvas);
}

function updateFrame(delta) {
	
}

function drawFrame() {
	
}
