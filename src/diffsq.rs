pub fn diff_square(n: i64) -> i64 {
    return square_sum(n) - sum_square(n);
}

fn square_sum(n: i64) -> i64 {
    let mut s = 0;
    for i in 1..n+1{
        s += i;
    }
    return s * s;
}

#[test]
fn test_square_sum(){
    assert_eq!(3025, square_sum(10));
}

fn sum_square(n: i64) -> i64 {
    let mut s = 0;
    for i in 1..n+1 {
        s += i * i;
    }
    return s;
}

#[test]
fn test_sum_square(){
    assert_eq!(385, sum_square(10));    
}
