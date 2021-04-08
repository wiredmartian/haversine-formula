## Calculate distance

Calculate distance between 2 coordinates on a sphere using the haversine formula


## Haversine Formula

``` 
a = sin²(φB - φA/2) + cos φA * cos φB * sin²(λB - λA/2)
c = 2 * atan2( √a, √(1−a) )
d = R ⋅ c
```


### Reference

> https://community.esri.com/t5/coordinate-reference-systems/distance-on-a-sphere-the-haversine-formula/ba-p/902128

> https://www.databasejournal.com/features/mysql/mysql-calculating-distance-based-on-latitude-and-longitude.html

> https://stackoverflow.com/a/27943/5483182
