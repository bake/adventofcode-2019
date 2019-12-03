<?php

declare(strict_types=1);

class Point implements \Ds\Hashable
{
	public int $x;
	public int $y;

	public function __construct(int $x = 0, int $y = 0)
	{
		$this->x = $x;
		$this->y = $y;
	}

	public function hash(): string
	{
		return $this->x . ' x ' . $this->y;
	}

	public function equals($obj): bool
	{
		return $this->x === $obj->x && $this->y === $obj->y;
	}
}

class Step
{
	public string $direction;
	public int $distance;

	public function __construct(string $input)
	{
		$this->direction = substr($input, 0, 1);
		$this->distance = (int) substr($input, 1);
	}
}

$inputs = file_get_contents('03.txt');
$inputs = explode(PHP_EOL, $inputs) ?: [];
$inputs = array_filter($inputs);
$inputs = array_map(function ($input) {
	$steps = explode(',', $input) ?: [];
	$steps = array_map((fn ($input) => new Step($input)), $steps);
	return $steps;
}, $inputs);

$maps = array_map((fn () => new \Ds\Map), $inputs);

$dx = ['U' => 0, 'D' => 0, 'L' => -1, 'R' => 1];
$dy = ['U' => -1, 'D' => 1, 'L' => 0, 'R' => 0];
foreach ($inputs as $i => $steps) {
	$k = 0;
	$location = new Point;
	foreach ($steps as $j => $step) {
		foreach (range(0, $step->distance - 1) as $_) {
			$location->x += $dx[$step->direction];
			$location->y += $dy[$step->direction];
			$maps[$i]->put(clone $location, ++$k);
		}
	}
}

$intersections = $maps[0]->intersect($maps[1]);
$intersections->ksort(function ($a, $b) {
	if ($a->x == 0 && $a->y == 0 && $b->x == 0 && $b->y == 0) {
		return INF;
	}
	return abs($a->x) + abs($a->y) <=> abs($b->x) + abs($b->y);
});
$point = $intersections->first()->key;

echo abs($point->x) + abs($point->y) . PHP_EOL;

$intersections->ksort(function ($a, $b) use ($maps) {
	$a = array_sum(array_map((fn ($map) => $map->get($a)), $maps));
	$b = array_sum(array_map((fn ($map) => $map->get($b)), $maps));
	return $a <=> $b;
});
$point = $intersections->first()->key;
echo array_sum(array_map((fn ($map) => $map->get($point)), $maps)) . PHP_EOL;
