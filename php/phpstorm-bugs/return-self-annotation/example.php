<?php
trait AzrTrait {
    /** @return self */
    public function setAzr(): self { return $this; }
}

trait BarTrait {
    /** @return self */
    public function setBar() { return $this; }
}

trait CraTrait {
    public function setCra(): self { return $this; }
}

trait DarTrait {
    /** @return $this */
    public function setDar(): self { return $this; }
}

trait EraTrait {
    /** @return static */
    public function setEra(): self { return $this; }
}

class FooClass {
    use AzrTrait;
    use BarTrait;
    use CraTrait;
    use DarTrait;
    use EraTrait;
}

$foo = new FooClass();
