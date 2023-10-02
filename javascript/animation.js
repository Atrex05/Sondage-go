// animation.js
import * as THREE from 'three';

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer();

renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

const particles = new THREE.Group();
scene.add(particles);

const particleGeometry = new THREE.SphereGeometry(0.1, 32, 32);
const particleMaterial = new THREE.MeshBasicMaterial({ color: 0xff0000 });

for (let i = 0; i < 1000; i++) {
    const particle = new THREE.Mesh(particleGeometry, particleMaterial);
    particle.position.set(
        (Math.random() - 0.5) * 10,
        (Math.random() - 0.5) * 10,
        (Math.random() - 0.5) * 10
    );
    particles.add(particle);
}

camera.position.z = 5;

const animate = () => {
    requestAnimationFrame(animate);

    particles.rotation.x += 0.005;
    particles.rotation.y += 0.005;

    renderer.render(scene, camera);
};

animate();
