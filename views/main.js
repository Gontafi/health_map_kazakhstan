$('path').on('click', function() {
    const jsonData = JSON.parse(document.getElementById("data").textContent);
    const metaJson = JSON.parse(document.getElementById("meta").textContent);
    console.log(jsonData);

    const regionId = $(this).attr('region-id');
    console.log(regionId);

    $('.grid_overlay').html('');
    console.log(jsonData);
    var sick = 0;
    var dead = 0;
    var cured = 0;

    for (const key in jsonData) {
        const data = jsonData[key];
        console.log(regionId, key);

        if (regionId === key) {
           sick = data.sick;
           dead = data.dead;
           cured = data.cured;
        }
    }

    $('.grid_overlay').append(`
    <div>
      <h2>${metaJson[regionId]}</h2>
      <p style="color: rgba(181,45,45,0.89)">Заболевших: ${sick}</p>
      <p style="color: rgba(0,0,0,0.71)">Смертей: ${dead}</p>
      <p style="color: rgba(68,136,17,0.89)">Выздоровевших: ${cured}</p>
    </div>
    `);
});

