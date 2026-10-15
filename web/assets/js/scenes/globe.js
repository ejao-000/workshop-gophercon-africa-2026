// The home page's 3D globe: a green sphere with every sighting plotted as a
// point at its coordinates. Reads the pairs from the #globe-points JSON.

import { mount } from '/assets/js/threejs.js';

mount('globe', function ({ THREE, scene, camera, colour, pointer }) {
	const group = new THREE.Group();
	scene.add(group);

	const radius = 2;
	group.add(
		new THREE.Mesh(
			new THREE.SphereGeometry(radius, 48, 32),
			new THREE.MeshStandardMaterial({
				color: colour('--three-primary', '#15803d'),
				roughness: 0.85,
				metalness: 0.05,
			}),
		),
	);

	group.add(
		new THREE.LineSegments(
			new THREE.WireframeGeometry(new THREE.SphereGeometry(radius + 0.02, 18, 12)),
			new THREE.LineBasicMaterial({
				color: colour('--three-accent', '#d97706'),
				transparent: true,
				opacity: 0.16,
			}),
		),
	);

	let coords = [];
	try {
		const el = document.getElementById('globe-points');
		coords = el ? JSON.parse(el.textContent) : [];
	} catch (err) {
		coords = [];
	}

	if (coords.length) {
		const positions = new Float32Array(coords.length * 3);
		const r = radius + 0.03;
		coords.forEach(function (p, i) {
			const lat = (p[0] * Math.PI) / 180;
			const lng = (p[1] * Math.PI) / 180;
			positions[i * 3] = r * Math.cos(lat) * Math.sin(lng);
			positions[i * 3 + 1] = r * Math.sin(lat);
			positions[i * 3 + 2] = r * Math.cos(lat) * Math.cos(lng);
		});
		const geometry = new THREE.BufferGeometry();
		geometry.setAttribute('position', new THREE.BufferAttribute(positions, 3));
		group.add(
			new THREE.Points(
				geometry,
				new THREE.PointsMaterial({
					color: colour('--three-accent', '#d97706'),
					size: 0.07,
					sizeAttenuation: true,
				}),
			),
		);
	}

	scene.add(new THREE.AmbientLight(0xffffff, 0.9));
	const key = new THREE.DirectionalLight(0xffffff, 2.2);
	key.position.set(4, 3, 5);
	scene.add(key);
	const rim = new THREE.DirectionalLight(colour('--three-accent', '#d97706'), 1.6);
	rim.position.set(-4, -2, -4);
	scene.add(rim);

	camera.position.z = 5.2;

	return function (time) {
		group.rotation.y = time * 0.12 + pointer.x * 0.35;
		group.rotation.x = pointer.y * 0.25;
	};
});
